package transaction

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	domain "github.com/leandrodam/transactions/internal/domain/transaction"
	"github.com/leandrodam/transactions/internal/infrastructure/exceptions"
)

func (h *handler) Create(c echo.Context) error {
	var req CreateTransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, exceptions.ErrBadRequest.ErrorJSON())
	}

	if err := c.Validate(&req); err != nil {
		e := exceptions.GetException(err).(*exceptions.Exception)
		return c.JSON(e.StatusCode, e.ErrorJSON())
	}

	transaction := domain.Transaction{
		AccountID:       req.AccountID,
		OperationTypeID: req.OperationTypeID,
		Amount:          req.Amount,
		EventDate:       time.Now().UTC(),
	}

	transaction, err := h.transactionUseCase.Create(c.Request().Context(), transaction)
	if err != nil {
		e := exceptions.GetException(err).(*exceptions.Exception)
		return c.JSON(e.StatusCode, e.ErrorJSON())
	}

	return c.JSON(http.StatusCreated, map[string]any{"data": TransactionResponse{
		TransactionID:   transaction.TransactionID,
		AccountID:       transaction.AccountID,
		OperationTypeID: transaction.OperationTypeID,
		Amount:          transaction.Amount,
		EventDate:       transaction.EventDate,
	}})
}
