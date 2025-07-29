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
		c.JSON(http.StatusBadRequest, exceptions.ErrBadRequest.ErrorJSON())
	}

	if err := c.Validate(&req); err != nil {
		e := exceptions.GetException(err).(*exceptions.Exception)
		return c.JSON(e.StatusCode, e.Messages)
	}

	transaction := domain.Transaction{
		AccountID:       req.AccountID,
		OperationTypeID: req.OperationTypeID,
		Amount:          req.Amount,
		EventDate:       time.Now().UTC(),
	}

	// TODO: Make request idempotent
	// Suggestion: Add an uuid field to the request payload.
	transaction, err := h.transactionUseCase.Create(c.Request().Context(), transaction)
	if err != nil {
		e := exceptions.GetException(err).(*exceptions.Exception)
		return c.JSON(e.StatusCode, e.Messages)
	}

	return c.JSON(http.StatusCreated, map[string]any{"data": transaction})
}
