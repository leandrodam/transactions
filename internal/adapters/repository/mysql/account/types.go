package account

import (
	"database/sql"

	domain "github.com/leandrodam/transactions/internal/domain/account"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.Repository {
	return &repository{
		db: db,
	}
}
