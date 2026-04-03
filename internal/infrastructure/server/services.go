package server

import (
	"github.com/leandrodam/transactions/internal/adapters/repository/mysql/account"
	"github.com/leandrodam/transactions/internal/adapters/repository/mysql/transaction"
	"github.com/leandrodam/transactions/internal/infrastructure"
	"github.com/leandrodam/transactions/internal/infrastructure/transactor"
)

func NewServices(resources infrastructure.Resources) infrastructure.Services {
	transactor, dbGetter := transactor.NewTransactor(resources.DB)

	return infrastructure.Services{
		Transactor:  transactor,
		Account:     account.NewRepository(dbGetter),
		Transaction: transaction.NewRepository(dbGetter),
	}
}
