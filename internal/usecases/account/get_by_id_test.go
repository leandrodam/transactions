package account

import (
	"context"
	"errors"
	"testing"

	domain "github.com/leandrodam/transactions/internal/domain/account"
	mocks "github.com/leandrodam/transactions/internal/domain/account/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_GetByID(t *testing.T) {
	tests := []struct {
		name              string
		input             int
		accountRepository *mocks.MockRepository
		output            domain.Account
		wantError         bool
	}{
		{
			name:  "Success",
			input: 1,
			accountRepository: func() *mocks.MockRepository {
				mockRepo := mocks.NewMockRepository(t)
				mockRepo.On("GetByID", mock.Anything, 1).
					Return(domain.Account{
						AccountID:      1,
						DocumentNumber: "12345678900",
					}, nil)
				return mockRepo
			}(),
			output: domain.Account{
				AccountID:      1,
				DocumentNumber: "12345678900",
			},
			wantError: false,
		},
		{
			name:  "Error fetching account by id",
			input: 1,
			accountRepository: func() *mocks.MockRepository {
				mockRepo := mocks.NewMockRepository(t)
				mockRepo.On("GetByID", mock.Anything, 1).
					Return(domain.Account{}, errors.New("exec error"))
				return mockRepo
			}(),
			output:    domain.Account{},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase := NewUseCase(tt.accountRepository)
			account, err := useCase.GetByID(context.Background(), tt.input)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.output, account)
		})
	}
}
