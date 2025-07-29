package transaction

import (
	"context"

	domain "github.com/leandrodam/transactions/internal/domain/transaction"
	"github.com/leandrodam/transactions/internal/infrastructure/exceptions"
)

func (r *repository) Create(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error) {
	query := `INSERT INTO transaction (account_id, operation_type_id, amount, event_date) VALUES (?,?,?,?)`

	args := []any{
		transaction.AccountID,
		transaction.OperationTypeID,
		transaction.Amount,
		transaction.EventDate,
	}

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return domain.Transaction{}, exceptions.GetException(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Transaction{}, exceptions.GetException(err)
	}

	transaction.TransactionID = int(id)

	return transaction, nil
}
