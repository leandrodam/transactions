package exceptions

import "net/http"

var (
	ErrInternal = &Exception{
		StatusCode: http.StatusInternalServerError,
		Messages:   []string{"Error detected while processing your request."},
	}

	ErrBadRequest = &Exception{
		StatusCode: http.StatusBadRequest,
		Messages:   []string{"The given values are invalid."},
	}

	ErrConflict = &Exception{
		StatusCode: http.StatusConflict,
		Messages:   []string{"The given values are already registered."},
	}

	ErrAccountNotFound = &Exception{
		StatusCode: http.StatusNotFound,
		Messages:   []string{"The requested account was not found."},
	}

	ErrInvalidAccountID = &Exception{
		StatusCode: http.StatusBadRequest,
		Messages:   []string{"The given accountID is invalid."},
	}

	ErrTransactionNotFound = &Exception{
		StatusCode: http.StatusNotFound,
		Messages:   []string{"The requested transaction was not found."},
	}
)
