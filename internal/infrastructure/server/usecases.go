package server

import (
	"github.com/leandrodam/transactions/internal/infrastructure"
	"github.com/leandrodam/transactions/internal/usecases/account"
	"github.com/leandrodam/transactions/internal/usecases/transaction"
)

func NewUseCases(services infrastructure.Services) infrastructure.UseCases {
	return infrastructure.UseCases{
		Account:     account.NewUseCase(services.Account),
		Transaction: transaction.NewUseCase(services.Transaction),
	}
}
