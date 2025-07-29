package account

import (
	"context"

	domain "github.com/leandrodam/transactions/internal/domain/account"
	"github.com/leandrodam/transactions/internal/infrastructure/exceptions"
)

func (r *repository) Create(ctx context.Context, account domain.Account) (domain.Account, error) {
	query := `INSERT INTO account (document_number) VALUES (?)`

	result, err := r.db.ExecContext(ctx, query, account.DocumentNumber)
	if err != nil {
		return domain.Account{}, exceptions.GetException(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Account{}, exceptions.GetException(err)
	}

	account.AccountID = int(id)

	return account, nil
}
