package account

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	domain "github.com/leandrodam/transactions/internal/domain/account"
	"github.com/leandrodam/transactions/internal/infrastructure/exceptions"
	"github.com/stretchr/testify/assert"
)

func Test_GetById(t *testing.T) {
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
				mock.ExpectQuery("SELECT account_id, document_number FROM account").
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
				mock.ExpectQuery("SELECT account_id, document_number FROM account").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"account_id", "document_number"}))
			},
			output:        domain.Account{},
			expectedError: exceptions.ErrAccountNotFound,
		},
		{
			name:  "Error rows scan",
			input: 1,
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT account_id, document_number FROM account").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"account_id", "document_number"}).
						AddRow("erro", "12345678900"))
			},
			output:        domain.Account{},
			expectedError: exceptions.ErrInternal,
		},
		{
			name:  "Success",
			input: 1,
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT account_id, document_number FROM account").
					WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{"account_id", "document_number"}).
							AddRow(1, "12345678900"),
					)
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

			repo := NewRepository(db)

			account, err := repo.GetByID(context.Background(), tt.input)
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.output, account)

			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}
