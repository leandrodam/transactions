package server

import (
	"github.com/leandrodam/transactions/internal/infrastructure"
	"github.com/leandrodam/transactions/internal/usecases/account"
	"github.com/leandrodam/transactions/internal/usecases/transaction"
)

func NewUseCases(services infrastructure.Services) infrastructure.UseCases {
	useCases := infrastructure.UseCases{}

	useCases.Account = account.NewUseCase(services.Account)
	useCases.Transaction = transaction.NewUseCase(services.Transaction, services.Account)

	return useCases
}
