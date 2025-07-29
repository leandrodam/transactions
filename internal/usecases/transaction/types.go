package transaction

import (
	"context"
	"errors"

	domain "github.com/leandrodam/transactions/internal/domain/transaction"
)

type UseCase interface {
	Create(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error)
}

type useCase struct {
	transactionRepository domain.Repository
}

func NewUseCase(transactionRepository domain.Repository) UseCase {
	return &useCase{
		transactionRepository: transactionRepository,
	}
}

var (
	ErrInvalidAccountID         = errors.New("invalid account id")
	ErrInvalidOperationTypeID   = errors.New("invalid operation type id")
	ErrInvalidTransactionAmount = errors.New("invalid transaction amount")
)
