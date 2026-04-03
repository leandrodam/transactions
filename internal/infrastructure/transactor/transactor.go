package transactor

import (
	"context"
	"database/sql"
	"errors"
)

type Transactor interface {
	WithinTransaction(context.Context, func(context.Context) error) error
}

type transactor struct {
	db *sql.DB
}

func NewTransactor(db *sql.DB) (Transactor, DBGetter) {
	return &transactor{db: db}, func(ctx context.Context) DB {
		if tx := fromContext(ctx); tx != nil {
			return tx
		}

		return db
	}
}

func (t *transactor) WithinTransaction(ctx context.Context, fn func(context.Context) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Join(ErrBeginTransaction, err)
	}

	txCtx := newContext(ctx, tx)

	if err := fn(txCtx); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return errors.Join(ErrRollbackTransaction, rollbackErr, err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return errors.Join(ErrCommitTransaction, err)
	}

	return nil
}

func IsWithinTransaction(ctx context.Context) bool {
	return ctx.Value(transactorKey{}) != nil
}
