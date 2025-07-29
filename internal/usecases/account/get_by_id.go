package account

import (
	"context"

	domain "github.com/leandrodam/transactions/internal/domain/account"
)

func (uc *useCase) GetByID(ctx context.Context, accountID int) (domain.Account, error) {
	return uc.accountRepository.GetByID(ctx, accountID)
}
