package account

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	domain "github.com/leandrodam/transactions/internal/domain/account"
	"github.com/leandrodam/transactions/internal/infrastructure/exceptions"
	"github.com/leandrodam/transactions/internal/infrastructure/transactor"
	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	tests := []struct {
		name          string
		input         domain.Account
		mockFunc      func(sqlmock.Sqlmock)
		output        domain.Account
		expectedError error
	}{
		{
			name: "Duplicate entry",
			input: domain.Account{
				DocumentNumber: "12345678900",
			},
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO account").
					WithArgs("12345678900").
					WillReturnError(&mysql.MySQLError{Number: 1062})
			},
			output:        domain.Account{},
			expectedError: exceptions.ErrConflict,
		},
		{
			name: "Error fetching last insert id",
			input: domain.Account{
				DocumentNumber: "12345678900",
			},
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO account").
					WithArgs("12345678900").
					WillReturnResult(sqlmock.NewErrorResult(exceptions.ErrInternal))
			},
			output:        domain.Account{},
			expectedError: exceptions.ErrInternal,
		},
		{
			name: "Success",
			input: domain.Account{
				DocumentNumber: "12345678900",
			},
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO account").
					WithArgs("12345678900").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			output: domain.Account{
				AccountID:      1,
				DocumentNumber: "12345678900",
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			if tt.mockFunc != nil {
				tt.mockFunc(mock)
			}

			_, dbGetter := transactor.NewTransactor(db)

			repo := NewRepository(dbGetter)

			account, err := repo.Create(context.Background(), tt.input)
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.output, account)

			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}

func Test_UpdateBalance(t *testing.T) {
	tests := []struct {
		name          string
		accountID     int
		amount        float64
		mockFunc      func(sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name:      "Internal error",
			accountID: 1,
			amount:    100.00,
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("UPDATE account SET available_credit = available_credit +").
					WithArgs(100.00, 1).
					WillReturnError(&mysql.MySQLError{Number: 1})
			},
			expectedError: exceptions.ErrInternal,
		},
		{
			name:      "Success",
			accountID: 1,
			amount:    100.00,
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("UPDATE account SET available_credit = available_credit +").
					WithArgs(100.00, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			if tt.mockFunc != nil {
				tt.mockFunc(mock)
			}

			_, dbGetter := transactor.NewTransactor(db)

			repo := NewRepository(dbGetter)

			err = repo.UpdateBalance(context.Background(), tt.accountID, tt.amount)
			assert.Equal(t, tt.expectedError, err)

			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}

func Test_Find(t *testing.T) {
	tests := []struct {
		name          string
		input         int
		mockFunc      func(sqlmock.Sqlmock)
		output        domain.Account
		expectedError error
	}{
		{
			name:  "Internal error",
			input: 1,
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT account_id, document_number, available_credit FROM account").
					WithArgs(1).
					WillReturnError(&mysql.MySQLError{Number: 1})
			},
			output:        domain.Account{},
			expectedError: exceptions.ErrInternal,
		},
		{
			name:  "Account not found",
			input: 1,
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT account_id, document_number, available_credit FROM account").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"account_id", "document_number", "available_credit"}))
			},
			output:        domain.Account{},
			expectedError: exceptions.ErrAccountNotFound,
		},
		{
			name:  "Error rows scan",
			input: 1,
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT account_id, document_number, available_credit FROM account").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"account_id", "document_number", "available_credit"}).
						AddRow("error", "12345678900", "0"))
			},
			output:        domain.Account{},
			expectedError: exceptions.ErrInternal,
		},
		{
			name:  "Success",
			input: 1,
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT account_id, document_number, available_credit FROM account").
					WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{"account_id", "document_number", "available_credit"}).
							AddRow(1, "12345678900", "300.00"),
					)
			},
			output: domain.Account{
				AccountID:       1,
				DocumentNumber:  "12345678900",
				AvailableCredit: 300.00,
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			if tt.mockFunc != nil {
				tt.mockFunc(mock)
			}

			_, dbGetter := transactor.NewTransactor(db)

			repo := NewRepository(dbGetter)

			account, err := repo.Find(context.Background(), tt.input)
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.output, account)

			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}
