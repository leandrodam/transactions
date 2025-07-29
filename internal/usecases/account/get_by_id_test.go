package account

import (
	"context"
	"errors"
	"testing"

	domain "github.com/leandrodam/transactions/internal/domain/account"
	mock "github.com/leandrodam/transactions/internal/domain/account/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_GetByID(t *testing.T) {
	tests := []struct {
		name              string
		accountID         int
		accountRepository *mock.MockRepository
		expectedAccount   domain.Account
		wantError         bool
	}{
		{
			name:      "Success",
			accountID: 1,
			accountRepository: func() *mock.MockRepository {
				mockRepo := mock.NewMockRepository(t)
				mockRepo.On("GetByID", context.Background(), 1).
					Return(domain.Account{
						AccountID:      1,
						DocumentNumber: "12345678900",
					}, nil)
				return mockRepo
			}(),
			expectedAccount: domain.Account{
				AccountID:      1,
				DocumentNumber: "12345678900",
			},
			wantError: false,
		},
		{
			name:      "Error fetching account by id",
			accountID: 1,
			accountRepository: func() *mock.MockRepository {
				mockRepo := mock.NewMockRepository(t)
				mockRepo.On("GetByID", context.Background(), 1).
					Return(domain.Account{}, errors.New("exec error"))
				return mockRepo
			}(),
			expectedAccount: domain.Account{},
			wantError:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase := NewUseCase(tt.accountRepository)
			account, err := useCase.GetByID(context.Background(), tt.accountID)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedAccount, account)
		})
	}
}
