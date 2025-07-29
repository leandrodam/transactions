package transaction

import (
	"database/sql"

	domain "github.com/leandrodam/transactions/internal/domain/transaction"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.Repository {
	return &repository{
		db: db,
	}
}
