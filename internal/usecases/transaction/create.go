package transaction

import (
	"context"

	domain "github.com/leandrodam/transactions/internal/domain/transaction"
)

func (uc *useCase) Create(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error) {
	transaction.AdjustAmountByOperationType()
	return uc.transactionRepository.Create(ctx, transaction)
}
