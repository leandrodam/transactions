package account

import (
	"context"

	domain "github.com/leandrodam/transactions/internal/domain/account"
	"github.com/shopspring/decimal"
)

func (uc *useCase) Create(ctx context.Context, account domain.Account) (domain.Account, error) {
	return uc.accountRepository.Create(ctx, account)
}

func (uc *useCase) Find(ctx context.Context, accountID int) (domain.Account, error) {
	return uc.accountRepository.Find(ctx, accountID)
}

func (uc *useCase) UpdateBalance(ctx context.Context, accountID int, amount decimal.Decimal) error {
	return uc.accountRepository.UpdateBalance(ctx, accountID, amount)
}
