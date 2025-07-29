package transaction

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, transaction Transaction) (Transaction, error)
}
