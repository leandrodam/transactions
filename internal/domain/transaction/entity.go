package transaction

import (
	"time"

	operationtypedomain "github.com/leandrodam/transactions/internal/domain/operationtype"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	TransactionID   int
	AccountID       int
	OperationTypeID int
	Amount          decimal.Decimal
	EventDate       time.Time
}

func (t *Transaction) AdjustAmountByOperationType() {
	switch t.OperationTypeID {
	case
		operationtypedomain.TypeNormalPurchase,
		operationtypedomain.TypePurchaseWithInstallments,
		operationtypedomain.TypeWithdrawal:
		t.Amount = t.Amount.Neg()
	}
}
