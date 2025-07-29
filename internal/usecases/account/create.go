package account

import (
	"context"
	"errors"

	domain "github.com/leandrodam/transactions/internal/domain/account"
)

var (
	ErrCreateAccount = errors.New("can't create account")
)

func (uc *useCase) Create(ctx context.Context, account domain.Account) (domain.Account, error) {
	return uc.accountRepository.Create(ctx, account)
}
