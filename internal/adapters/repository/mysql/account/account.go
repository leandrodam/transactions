package account

import (
	"context"

	domain "github.com/leandrodam/transactions/internal/domain/account"
	"github.com/leandrodam/transactions/internal/infrastructure/exceptions"
)

func (r *repository) Create(ctx context.Context, account domain.Account) (domain.Account, error) {
	query := `INSERT INTO account (document_number) VALUES (?)`

	result, err := r.dbGetter(ctx).ExecContext(ctx, query, account.DocumentNumber)
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

func (r *repository) UpdateBalance(ctx context.Context, accountID int, amount float64) error {
	query := `UPDATE account SET available_credit = available_credit + ? WHERE account_id = ?`

	_, err := r.dbGetter(ctx).ExecContext(ctx, query, amount, accountID)
	if err != nil {
		return exceptions.GetException(err)
	}

	return nil
}

func (r *repository) Find(ctx context.Context, id int) (domain.Account, error) {
	query := `SELECT account_id, document_number, available_credit FROM account WHERE account_id = ?`

	rows, err := r.dbGetter(ctx).QueryContext(ctx, query, id)
	if err != nil {
		return domain.Account{}, exceptions.GetException(err)
	}

	defer rows.Close()

	if !rows.Next() {
		return domain.Account{}, exceptions.ErrAccountNotFound
	}

	account := domain.Account{}
	err = rows.Scan(&account.AccountID, &account.DocumentNumber, &account.AvailableCredit)
	if err != nil {
		return domain.Account{}, exceptions.GetException(err)
	}

	return account, nil
}
