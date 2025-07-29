package account

import (
	"net/http"

	"github.com/labstack/echo/v4"
	domain "github.com/leandrodam/transactions/internal/domain/account"
	"github.com/leandrodam/transactions/internal/infrastructure/exceptions"
)

func (h *handler) Create(c echo.Context) error {
	var req CreateAccountRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, exceptions.ErrBadRequest.ErrorJSON())
	}

	if err := c.Validate(&req); err != nil {
		e := exceptions.GetException(err).(*exceptions.Exception)
		c.JSON(http.StatusBadRequest, e.ErrorJSON())
	}

	account := domain.Account{
		DocumentNumber: req.DocumentNumber,
	}

	account, err := h.accountUseCase.Create(c.Request().Context(), account)
	if err != nil {
		e := exceptions.GetException(err).(*exceptions.Exception)
		return c.JSON(e.StatusCode, e.ErrorJSON())
	}

	return c.JSON(http.StatusCreated, map[string]any{"data": account})
}
