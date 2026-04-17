package transaction

import (
	"time"

	"github.com/labstack/echo/v4"
	usecase "github.com/leandrodam/transactions/internal/usecases/transaction"
	"github.com/shopspring/decimal"
)

type (
	Handler interface {
		Create(c echo.Context) error
	}

	handler struct {
		transactionUseCase usecase.UseCase
	}

	CreateTransactionRequest struct {
		AccountID       int             `json:"account_id" validate:"required,gt=0"`
		OperationTypeID int             `json:"operation_type_id" validate:"required,gt=0"`
		Amount          decimal.Decimal `json:"amount" validate:"dgte=0"`
	}

	TransactionResponse struct {
		TransactionID   int             `json:"transaction_id"`
		AccountID       int             `json:"account_id,omitempty"`
		OperationTypeID int             `json:"operation_type_id,omitempty"`
		Amount          decimal.Decimal `json:"amount,omitzero"`
		EventDate       time.Time       `json:"event_date,omitzero"`
	}
)

func NewHandler(useCase usecase.UseCase) Handler {
	return &handler{
		transactionUseCase: useCase,
	}
}
