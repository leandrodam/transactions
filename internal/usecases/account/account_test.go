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

func Test_Create(t *testing.T) {
	tests := []struct {
		name              string
		input             domain.Account
		accountRepository *mocks.MockRepository
		output            domain.Account
		wantError         bool
	}{
		{
			name: "Success",
			input: domain.Account{
				DocumentNumber: "12345678900",
			},
			accountRepository: func() *mocks.MockRepository {
				m := mocks.NewMockRepository(t)
				m.On("Create", mock.Anything, domain.Account{
					DocumentNumber: "12345678900",
				}).Return(domain.Account{
					AccountID:      1,
					DocumentNumber: "12345678900",
				}, nil)
				return m
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
			accountRepository: func() *mocks.MockRepository {
				m := mocks.NewMockRepository(t)
				m.On("Create", mock.Anything, domain.Account{
					DocumentNumber: "12345678900",
				}).Return(domain.Account{}, errors.New("exec error"))
				return m
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

func Test_Find(t *testing.T) {
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
				m := mocks.NewMockRepository(t)
				m.On("Find", mock.Anything, 1).
					Return(domain.Account{
						AccountID:      1,
						DocumentNumber: "12345678900",
					}, nil)
				return m
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
				m := mocks.NewMockRepository(t)
				m.On("Find", mock.Anything, 1).
					Return(domain.Account{}, errors.New("exec error"))
				return m
			}(),
			output:    domain.Account{},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase := NewUseCase(tt.accountRepository)
			account, err := useCase.Find(context.Background(), tt.input)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.output, account)
		})
	}
}
