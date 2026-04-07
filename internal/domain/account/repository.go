package account

import (
	"context"

	"github.com/shopspring/decimal"
)

type Repository interface {
	Create(ctx context.Context, account Account) (Account, error)
	Find(ctx context.Context, id int) (Account, error)
	UpdateBalance(ctx context.Context, accountID int, amount decimal.Decimal) error
}
