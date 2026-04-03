package transaction

import (
	domain "github.com/leandrodam/transactions/internal/domain/transaction"
	"github.com/leandrodam/transactions/internal/infrastructure/transactor"
)

type repository struct {
	dbGetter transactor.DBGetter
}

func NewRepository(dbGetter transactor.DBGetter) domain.Repository {
	return &repository{dbGetter: dbGetter}
}
