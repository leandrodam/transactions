package transaction

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	domain "github.com/leandrodam/transactions/internal/domain/transaction"
	"github.com/leandrodam/transactions/internal/infrastructure/exceptions"
	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	eventDate := time.Now().UTC()

	tests := []struct {
		name          string
		input         domain.Transaction
		mockFunc      func(sqlmock.Sqlmock)
		output        domain.Transaction
		expectedError error
	}{
		{
			name: "Duplicate entry",
			input: domain.Transaction{
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          123.45,
				EventDate:       eventDate,
			},
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO transaction").
					WithArgs(1, 1, 123.45, eventDate).
					WillReturnError(&mysql.MySQLError{Number: 1062})
			},
			output:        domain.Transaction{},
			expectedError: exceptions.ErrConflict,
		},
		{
			name: "Error fetching last insert id",
			input: domain.Transaction{
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          123.45,
				EventDate:       eventDate,
			},
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO transaction").
					WithArgs(1, 1, 123.45, eventDate).
					WillReturnResult(sqlmock.NewErrorResult(exceptions.ErrInternal))
			},
			output:        domain.Transaction{},
			expectedError: exceptions.ErrInternal,
		},
		{
			name: "Success",
			input: domain.Transaction{
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          123.45,
				EventDate:       eventDate,
			},
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO transaction").
					WithArgs(1, 1, 123.45, eventDate).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			output: domain.Transaction{
				TransactionID:   1,
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          123.45,
				EventDate:       eventDate,
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

			transaction, err := repo.Create(context.Background(), tt.input)
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.output, transaction)

			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}
