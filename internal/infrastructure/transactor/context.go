package transactor

import (
	"context"
	"database/sql"
)

type transactorKey struct{}

func newContext(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, transactorKey{}, tx)
}

func fromContext(ctx context.Context) DB {
	if tx, ok := ctx.Value(transactorKey{}).(*sql.Tx); ok {
		return tx
	}
	return nil
}
