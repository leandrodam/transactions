package account

import (
	"github.com/labstack/echo/v4"
	usecase "github.com/leandrodam/transactions/internal/usecases/account"
)

type (
	Handler interface {
		Create(c echo.Context) error
		GetByID(c echo.Context) error
	}

	handler struct {
		accountUseCase usecase.UseCase
	}

	CreateAccountRequest struct {
		DocumentNumber string `json:"document_number" validate:"required,len=11,numeric"`
	}

	AccountResponse struct {
		AccountID int `json:"account_id"`
	}
)

func NewHandler(useCase usecase.UseCase) Handler {
	return &handler{
		accountUseCase: useCase,
	}
}
