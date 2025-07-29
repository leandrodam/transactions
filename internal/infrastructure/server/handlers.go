package server

import (
	"github.com/leandrodam/transactions/internal/adapters/http/handlers/account"
	"github.com/leandrodam/transactions/internal/adapters/http/handlers/transaction"
	"github.com/leandrodam/transactions/internal/infrastructure"
)

func NewHandlers(useCases infrastructure.UseCases) infrastructure.Handlers {
	return infrastructure.Handlers{
		Account:     account.NewHandler(useCases.Account),
		Transaction: transaction.NewHandler(useCases.Transaction),
	}
}
