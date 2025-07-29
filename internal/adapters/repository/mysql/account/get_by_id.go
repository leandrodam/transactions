package account

import (
	"context"

	domain "github.com/leandrodam/transactions/internal/domain/account"
	"github.com/leandrodam/transactions/internal/infrastructure/exceptions"
)

func (r *repository) GetByID(ctx context.Context, id int) (domain.Account, error) {
	query := `SELECT account_id, document_number FROM account WHERE account_id = ?`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return domain.Account{}, exceptions.GetException(err)
	}

	defer rows.Close()

	if !rows.Next() {
		return domain.Account{}, exceptions.ErrAccountNotFound
	}

	account := domain.Account{}
	err = rows.Scan(&account.AccountID, &account.DocumentNumber)
	if err != nil {
		return domain.Account{}, exceptions.GetException(err)
	}

	return account, nil
}
