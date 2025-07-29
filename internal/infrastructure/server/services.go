package server

import (
	"github.com/leandrodam/transactions/internal/adapters/repository/mysql/account"
	"github.com/leandrodam/transactions/internal/adapters/repository/mysql/transaction"
	"github.com/leandrodam/transactions/internal/infrastructure"
)

func NewServices(resources infrastructure.Resources) infrastructure.Services {
	return infrastructure.Services{
		Account:     account.NewRepository(resources.DB),
		Transaction: transaction.NewRepository(resources.DB),
	}
}
