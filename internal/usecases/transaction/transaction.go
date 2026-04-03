package transaction

import (
	"context"
	"errors"

	domain "github.com/leandrodam/transactions/internal/domain/transaction"
)

func (uc *useCase) Create(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error) {
	err := uc.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		transaction.AdjustAmountByOperationType()

		account, err := uc.accountRepository.Find(ctx, transaction.AccountID)
		if err != nil {
			return err
		}

		account.AvailableCredit += transaction.Amount
		if account.AvailableCredit < 0 {
			return errors.New("negative balance")
		}

		err = uc.accountRepository.UpdateBalance(ctx, account.AccountID, transaction.Amount)
		if err != nil {
			return err
		}

		transaction, err = uc.transactionRepository.Create(ctx, transaction)

		return err
	})

	if err != nil {
		return domain.Transaction{}, err
	}

	return transaction, nil
}
