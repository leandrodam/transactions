package transaction

import (
	"time"

	operationtypedomain "github.com/leandrodam/transactions/internal/domain/operationtype"
)

type Transaction struct {
	TransactionID   int
	AccountID       int
	OperationTypeID int
	Amount          float64
	EventDate       time.Time
}

func (t *Transaction) AdjustAmountByOperationType() {
	switch t.OperationTypeID {
	case
		operationtypedomain.TypeNormalPurchase,
		operationtypedomain.TypePurchaseWithInstallments,
		operationtypedomain.TypeWithdrawal:
		t.Amount = -1 * t.Amount
	}
}
