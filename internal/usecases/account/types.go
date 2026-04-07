package account

import (
	"context"

	domain "github.com/leandrodam/transactions/internal/domain/account"
	"github.com/shopspring/decimal"
)

type UseCase interface {
	Create(ctx context.Context, account domain.Account) (domain.Account, error)
	UpdateBalance(ctx context.Context, accountID int, amount decimal.Decimal) error
	Find(ctx context.Context, accountID int) (domain.Account, error)
}

type useCase struct {
	accountRepository domain.Repository
}

func NewUseCase(accountRepository domain.Repository) UseCase {
	return &useCase{
		accountRepository: accountRepository,
	}
}
