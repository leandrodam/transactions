package transaction

import (
	"context"
	"errors"

	domain "github.com/leandrodam/transactions/internal/domain/transaction"
)

func (uc *useCase) Create(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error) {
	transaction.AdjustAmountByOperationType()

	account, err := uc.accountRepository.Find(ctx, transaction.AccountID)
	if err != nil {
		return domain.Transaction{}, err
	}

	account.AvailableCredit += transaction.Amount
	if account.AvailableCredit < 0 {
		return domain.Transaction{}, errors.New("negative balance")
	}

	err = uc.accountRepository.UpdateBalance(ctx, account.AccountID, transaction.Amount)
	if err != nil {
		return domain.Transaction{}, err
	}

	return uc.transactionRepository.Create(ctx, transaction)
}
