package account

import "github.com/shopspring/decimal"

type Account struct {
	AccountID       int
	DocumentNumber  string
	AvailableCredit decimal.Decimal
}
