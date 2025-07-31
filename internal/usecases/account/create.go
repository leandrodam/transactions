package account

import (
	"context"

	domain "github.com/leandrodam/transactions/internal/domain/account"
)

func (uc *useCase) Create(ctx context.Context, account domain.Account) (domain.Account, error) {
	return uc.accountRepository.Create(ctx, account)
}
