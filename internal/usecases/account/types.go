package account

import (
	"context"

	domain "github.com/leandrodam/transactions/internal/domain/account"
)

type UseCase interface {
	Create(ctx context.Context, account domain.Account) (domain.Account, error)
	GetByID(ctx context.Context, accountID int) (domain.Account, error)
}

type useCase struct {
	accountRepository domain.Repository
}

func NewUseCase(accountRepository domain.Repository) UseCase {
	return &useCase{
		accountRepository: accountRepository,
	}
}
