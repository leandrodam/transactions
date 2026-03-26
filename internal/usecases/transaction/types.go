package transaction

import (
	"context"

	accountdomain "github.com/leandrodam/transactions/internal/domain/account"
	transactiondomain "github.com/leandrodam/transactions/internal/domain/transaction"
)

type UseCase interface {
	Create(ctx context.Context, transaction transactiondomain.Transaction) (transactiondomain.Transaction, error)
}

type useCase struct {
	transactionRepository transactiondomain.Repository
	accountRepository     accountdomain.Repository
}

func NewUseCase(transactionRepository transactiondomain.Repository, accountRepository accountdomain.Repository) UseCase {
	return &useCase{
		transactionRepository: transactionRepository,
		accountRepository:     accountRepository,
	}
}
