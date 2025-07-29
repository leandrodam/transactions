package exceptions

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/leandrodam/transactions/internal/infrastructure/validator"
)

type Exception struct {
	StatusCode int
	Messages   []string
}

func (e *Exception) Error() string {
	return strings.Join(e.Messages, ". ")
}

func (e *Exception) ErrorJSON() map[string]any {
	return map[string]any{
		"errors": e.Messages,
	}
}

func GetException(err error) error {
	if e, ok := err.(*Exception); ok {
		return e
	}

	if e, ok := err.(*mysql.MySQLError); ok {
		return getMySQLException(e)
	}

	if e, ok := err.(*validator.ValidationError); ok {
		return getValidationException(e)
	}

	return ErrInternal
}

func getMySQLException(err error) error {
	e := ErrInternal

	if mysqlError, ok := err.(*mysql.MySQLError); ok {
		switch mysqlError.Number {
		case 1452:
			e = ErrBadRequest

		case 1062:
			e = ErrConflict
		}
	}

	return e
}

func getValidationException(err error) error {
	e := ErrBadRequest

	if validationErrors, ok := err.(*validator.ValidationError); ok {
		e.Messages = validationErrors.GetMessages()
	}

	return e
}
