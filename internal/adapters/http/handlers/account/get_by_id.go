package account

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/leandrodam/transactions/internal/infrastructure/exceptions"
)

func (h *handler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.ErrInvalidAccountID.ErrorJSON())
	}

	account, err := h.accountUseCase.GetByID(c.Request().Context(), id)
	if err != nil {
		e := exceptions.GetException(err).(*exceptions.Exception)
		return c.JSON(e.StatusCode, e.ErrorJSON())
	}

	return c.JSON(http.StatusOK, map[string]any{"data": account})
}
