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

func Test_Create(t *testing.T) {
	tests := []struct {
		name           string
		input          domain.Account
		mockFunc       func(sqlmock.Sqlmock)
		expectedOutput domain.Account
		expectedError  error
	}{
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
			expectedOutput: domain.Account{
				AccountID:      1,
				DocumentNumber: "12345678900",
			},
			expectedError: nil,
		},
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
			expectedOutput: domain.Account{},
			expectedError:  exceptions.ErrConflict,
		},
		{
			name: "Internal error",
			input: domain.Account{
				DocumentNumber: "12345678900",
			},
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO account").
					WithArgs("12345678900").
					WillReturnError(&mysql.MySQLError{Number: 1})
			},
			expectedOutput: domain.Account{},
			expectedError:  exceptions.ErrInternal,
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

			account, err := repo.Create(context.Background(), tt.input)
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedOutput, account)

			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}
