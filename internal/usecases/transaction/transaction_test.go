package transaction

import (
	"context"
	"errors"
	"testing"
	"time"

	accountdomain "github.com/leandrodam/transactions/internal/domain/account"
	accountmocks "github.com/leandrodam/transactions/internal/domain/account/mocks"
	domain "github.com/leandrodam/transactions/internal/domain/transaction"
	mocks "github.com/leandrodam/transactions/internal/domain/transaction/mocks"
	"github.com/leandrodam/transactions/internal/infrastructure/transactor"
	transactormocks "github.com/leandrodam/transactions/internal/infrastructure/transactor/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Create(t *testing.T) {
	ctx := context.Background()

	eventDate := time.Now().UTC()

	successfulTransactor := func() *transactormocks.MockTransactor {
		m := transactormocks.NewMockTransactor(t)
		m.On("WithinTransaction", ctx, mock.Anything).
			Run(func(args mock.Arguments) {
				fn := args.Get(1).(func(context.Context) error)
				fn(ctx)
			}).
			Return(nil)
		return m
	}()

	failedTransactor := func() *transactormocks.MockTransactor {
		m := transactormocks.NewMockTransactor(t)
		m.On("WithinTransaction", ctx, mock.Anything).
			Run(func(args mock.Arguments) {
				fn := args.Get(1).(func(context.Context) error)
				fn(ctx)
			}).
			Return(errors.New("failed"))
		return m
	}()

	tests := []struct {
		name                  string
		input                 domain.Transaction
		transactionRepository *mocks.MockRepository
		accountRepository     *accountmocks.MockRepository
		transactor            transactor.Transactor
		output                domain.Transaction
		wantError             bool
	}{
		{
			name: "Error finding account",
			input: domain.Transaction{
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          123.45,
				EventDate:       eventDate,
			},
			accountRepository: func() *accountmocks.MockRepository {
				m := accountmocks.NewMockRepository(t)
				m.On("Find", ctx, 1).
					Return(accountdomain.Account{}, errors.New("failed"))
				return m
			}(),
			transactor: failedTransactor,
			output:     domain.Transaction{},
			wantError:  true,
		},
		{
			name: "Error: negative balance",
			input: domain.Transaction{
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          123.45,
				EventDate:       eventDate,
			},
			accountRepository: func() *accountmocks.MockRepository {
				m := accountmocks.NewMockRepository(t)
				m.On("Find", ctx, 1).
					Return(accountdomain.Account{
						AccountID:       1,
						DocumentNumber:  "12345",
						AvailableCredit: 100,
					}, nil)
				return m
			}(),
			transactor: failedTransactor,
			output:     domain.Transaction{},
			wantError:  true,
		},
		{
			name: "Error updating balance",
			input: domain.Transaction{
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          123.45,
				EventDate:       eventDate,
			},
			accountRepository: func() *accountmocks.MockRepository {
				m := accountmocks.NewMockRepository(t)
				m.On("Find", ctx, 1).
					Return(accountdomain.Account{
						AccountID:       1,
						DocumentNumber:  "12345",
						AvailableCredit: 1000,
					}, nil)
				m.On("UpdateBalance", ctx, 1, -123.45).
					Return(errors.New("failed"))
				return m
			}(),
			transactor: failedTransactor,
			output:     domain.Transaction{},
			wantError:  true,
		},
		{
			name: "Error creating transaction",
			input: domain.Transaction{
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          123.45,
				EventDate:       eventDate,
			},
			accountRepository: func() *accountmocks.MockRepository {
				m := accountmocks.NewMockRepository(t)
				m.On("Find", ctx, 1).
					Return(accountdomain.Account{
						AccountID:       1,
						DocumentNumber:  "12345",
						AvailableCredit: 1000,
					}, nil)
				m.On("UpdateBalance", ctx, 1, -123.45).
					Return(nil)
				return m
			}(),
			transactionRepository: func() *mocks.MockRepository {
				m := mocks.NewMockRepository(t)
				m.On("Create", ctx, mock.Anything).
					Return(domain.Transaction{}, errors.New("failed"))
				return m
			}(),
			transactor: failedTransactor,
			output:     domain.Transaction{},
			wantError:  true,
		},
		{
			name: "Successfully created transaction",
			input: domain.Transaction{
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          123.45,
				EventDate:       eventDate,
			},
			accountRepository: func() *accountmocks.MockRepository {
				m := accountmocks.NewMockRepository(t)
				m.On("Find", ctx, 1).
					Return(accountdomain.Account{
						AccountID:       1,
						DocumentNumber:  "12345",
						AvailableCredit: 1000,
					}, nil)
				m.On("UpdateBalance", ctx, 1, -123.45).
					Return(nil)
				return m
			}(),
			transactionRepository: func() *mocks.MockRepository {
				m := mocks.NewMockRepository(t)
				m.On("Create", ctx, mock.Anything).
					Return(domain.Transaction{
						TransactionID:   1,
						AccountID:       1,
						OperationTypeID: 1,
						Amount:          -123.45,
						EventDate:       eventDate,
					}, nil)
				return m
			}(),
			transactor: successfulTransactor,
			output: domain.Transaction{
				TransactionID:   1,
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          -123.45,
				EventDate:       eventDate,
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase := NewUseCase(tt.transactionRepository, tt.accountRepository, tt.transactor)
			transaction, err := useCase.Create(ctx, tt.input)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.output, transaction)
		})
	}
}
