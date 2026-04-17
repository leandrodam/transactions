package transaction

import (
	"context"

	domain "github.com/leandrodam/transactions/internal/domain/transaction"
	"github.com/leandrodam/transactions/internal/infrastructure/exceptions"
)

func (uc *useCase) Create(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error) {
	err := uc.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		transaction.AdjustAmountByOperationType()

		account, err := uc.accountRepository.Find(ctx, transaction.AccountID)
		if err != nil {
			return err
		}

		account.AvailableCredit = account.AvailableCredit.Add(transaction.Amount)
		if account.AvailableCredit.IsNegative() {
			return exceptions.ErrNegativeBalance
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
