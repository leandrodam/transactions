package transactor

import "errors"

var (
	ErrBeginTransaction    = errors.New("failed to begin transaction")
	ErrRollbackTransaction = errors.New("failed to rollback transaction")
	ErrCommitTransaction   = errors.New("failed to commit transaction")
)
