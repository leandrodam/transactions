package account

import (
	"context"
	"errors"
	"testing"

	domain "github.com/leandrodam/transactions/internal/domain/account"
	mock "github.com/leandrodam/transactions/internal/domain/account/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	tests := []struct {
		name              string
		input             domain.Account
		accountRepository *mock.MockRepository
		output            domain.Account
		wantError         bool
	}{
		{
			name: "Success",
			input: domain.Account{
				DocumentNumber: "12345678900",
			},
			accountRepository: func() *mock.MockRepository {
				mockRepo := mock.NewMockRepository(t)
				mockRepo.On("Create", context.Background(), domain.Account{
					DocumentNumber: "12345678900",
				}).Return(domain.Account{
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
			name: "Error creating account",
			input: domain.Account{
				DocumentNumber: "12345678900",
			},
			accountRepository: func() *mock.MockRepository {
				mockRepo := mock.NewMockRepository(t)
				mockRepo.On("Create", context.Background(), domain.Account{
					DocumentNumber: "12345678900",
				}).Return(domain.Account{}, errors.New("exec error"))
				return mockRepo
			}(),
			output:    domain.Account{},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase := NewUseCase(tt.accountRepository)
			account, err := useCase.Create(context.Background(), tt.input)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.output, account)
		})
	}
}
