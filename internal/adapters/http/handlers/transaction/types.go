package transaction

import (
	"github.com/labstack/echo/v4"
	usecase "github.com/leandrodam/transactions/internal/usecases/transaction"
)

type (
	Handler interface {
		Create(c echo.Context) error
	}

	handler struct {
		transactionUseCase usecase.UseCase
	}

	CreateTransactionRequest struct {
		AccountID       int     `json:"account_id" validate:"required,gt=0"`
		OperationTypeID int     `json:"operation_type_id" validate:"required,gt=0"`
		Amount          float64 `json:"amount" validate:"required"`
	}
)

func NewHandler(useCase usecase.UseCase) Handler {
	return &handler{
		transactionUseCase: useCase,
	}
}
