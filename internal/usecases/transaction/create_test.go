package transaction

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "github.com/leandrodam/transactions/internal/domain/transaction"
	mocks "github.com/leandrodam/transactions/internal/domain/transaction/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Create(t *testing.T) {
	eventDate := time.Now().UTC()

	tests := []struct {
		name                  string
		input                 domain.Transaction
		transactionRepository *mocks.MockRepository
		output                domain.Transaction
		wantError             bool
	}{
		{
			name: "Error creating transaction",
			input: domain.Transaction{
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          123.45,
				EventDate:       eventDate,
			},
			transactionRepository: func() *mocks.MockRepository {
				mockRepo := mocks.NewMockRepository(t)
				mockRepo.On("Create", mock.Anything, mock.MatchedBy(func(txn domain.Transaction) bool {
					return txn.AccountID == 1 &&
						txn.OperationTypeID == 1 &&
						txn.Amount == -123.45 &&
						time.Since(txn.EventDate) < 2*time.Second
				})).Return(domain.Transaction{}, errors.New("exec error"))
				return mockRepo
			}(),
			output:    domain.Transaction{},
			wantError: true,
		},
		{
			name: "Success: negative transaction",
			input: domain.Transaction{
				AccountID:       1,
				OperationTypeID: 4,
				Amount:          123.45,
				EventDate:       eventDate,
			},
			transactionRepository: func() *mocks.MockRepository {
				mockRepo := mocks.NewMockRepository(t)
				mockRepo.On("Create", mock.Anything, mock.MatchedBy(func(txn domain.Transaction) bool {
					return txn.AccountID == 1 &&
						txn.OperationTypeID == 4 &&
						txn.Amount == 123.45 &&
						time.Since(txn.EventDate) < 2*time.Second
				})).Return(domain.Transaction{
					TransactionID:   1,
					AccountID:       1,
					OperationTypeID: 4,
					Amount:          123.45,
					EventDate:       eventDate,
				}, nil)
				return mockRepo
			}(),
			output: domain.Transaction{
				TransactionID:   1,
				AccountID:       1,
				OperationTypeID: 4,
				Amount:          123.45,
				EventDate:       eventDate,
			},
			wantError: false,
		},
		{
			name: "Success: positive transaction",
			input: domain.Transaction{
				AccountID:       1,
				OperationTypeID: 4,
				Amount:          123.45,
				EventDate:       eventDate,
			},
			transactionRepository: func() *mocks.MockRepository {
				mockRepo := mocks.NewMockRepository(t)
				mockRepo.On("Create", mock.Anything, mock.MatchedBy(func(txn domain.Transaction) bool {
					return txn.AccountID == 1 &&
						txn.OperationTypeID == 4 &&
						txn.Amount == 123.45 &&
						time.Since(txn.EventDate) < 2*time.Second
				})).Return(domain.Transaction{
					TransactionID:   1,
					AccountID:       1,
					OperationTypeID: 4,
					Amount:          123.45,
					EventDate:       eventDate,
				}, nil)
				return mockRepo
			}(),
			output: domain.Transaction{
				TransactionID:   1,
				AccountID:       1,
				OperationTypeID: 4,
				Amount:          123.45,
				EventDate:       eventDate,
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase := NewUseCase(tt.transactionRepository)
			transaction, err := useCase.Create(context.Background(), tt.input)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.output, transaction)
		})
	}
}
