package account

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, account Account) (Account, error)
	GetByID(ctx context.Context, id int) (Account, error)
}
