package operationtype

type OperationType struct {
	OperationTypeID int
	Description     string
}

const (
	TypeNormalPurchase           = 1
	TypePurchaseWithInstallments = 2
	TypeWithdrawal               = 3
	TypeCreditVoucher            = 4
)
