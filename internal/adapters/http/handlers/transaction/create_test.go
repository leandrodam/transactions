package transaction

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	domain "github.com/leandrodam/transactions/internal/domain/transaction"
	"github.com/leandrodam/transactions/internal/infrastructure/exceptions"
	"github.com/leandrodam/transactions/internal/infrastructure/validator"
	mocks "github.com/leandrodam/transactions/internal/usecases/transaction/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Create(t *testing.T) {
	type output struct {
		statusCode int
		messages   map[string]any
	}

	eventDate := time.Now().UTC()

	tests := []struct {
		name    string
		input   []byte
		useCase *mocks.MockUseCase
		output  output
	}{
		{
			name: "Error binding parameters",
			input: []byte(`
				{
					"account_id": "1",
					"operation_type_id": "1",
					"amount": "123.45"
				}
			`),
			output: output{
				statusCode: http.StatusBadRequest,
				messages:   exceptions.ErrBadRequest.ErrorJSON(),
			},
		},
		{
			name: "Error: account id is required",
			input: []byte(`
				{
					"account_id": null,
					"operation_type_id": 1,
					"amount": 123.45
				}
			`),
			output: output{
				statusCode: http.StatusBadRequest,
				messages:   map[string]any{"errors": []string{"AccountID is required."}},
			},
		},
		{
			name: "Error: account id must be gt=0",
			input: []byte(`
				{
					"account_id": -1,
					"operation_type_id": 1,
					"amount": 123.45
				}
			`),
			output: output{
				statusCode: http.StatusBadRequest,
				messages:   map[string]any{"errors": []string{"AccountID must be greater than 0."}},
			},
		},
		{
			name: "Error: operation type id is required",
			input: []byte(`
				{
					"account_id": 1,
					"operation_type_id": null,
					"amount": 123.45
				}
			`),
			output: output{
				statusCode: http.StatusBadRequest,
				messages:   map[string]any{"errors": []string{"OperationTypeID is required."}},
			},
		},
		{
			name: "Error: operation type id must be gt=0",
			input: []byte(`
				{
					"account_id": 1,
					"operation_type_id": -1,
					"amount": 123.45
				}
			`),
			output: output{
				statusCode: http.StatusBadRequest,
				messages:   map[string]any{"errors": []string{"OperationTypeID must be greater than 0."}},
			},
		},
		{
			name: "Error: amount is required",
			input: []byte(`
				{
					"account_id": 1,
					"operation_type_id": 1,
					"amount": null
				}
			`),
			output: output{
				statusCode: http.StatusBadRequest,
				messages:   map[string]any{"errors": []string{"Amount is required."}},
			},
		},
		{
			name: "Error: amount must be gte=0",
			input: []byte(`
				{
					"account_id": 1,
					"operation_type_id": 1,
					"amount": -10.00
				}
			`),
			output: output{
				statusCode: http.StatusBadRequest,
				messages:   map[string]any{"errors": []string{"Amount must be greater than or equal to 0."}},
			},
		},
		{
			name: "Error creating transaction",
			input: []byte(`
				{
					"account_id": 1,
					"operation_type_id": 1,
					"amount": 123.45
				}
			`),
			useCase: func() *mocks.MockUseCase {
				mockUseCase := mocks.NewMockUseCase(t)
				mockUseCase.On("Create", mock.Anything, mock.MatchedBy(func(txn domain.Transaction) bool {
					return txn.AccountID == 1 &&
						txn.OperationTypeID == 1 &&
						txn.Amount == 123.45 &&
						time.Since(txn.EventDate) < 2*time.Second
				})).Return(domain.Transaction{}, exceptions.ErrInternal)
				return mockUseCase
			}(),
			output: output{
				statusCode: http.StatusInternalServerError,
				messages:   exceptions.ErrInternal.ErrorJSON(),
			},
		},
		{
			name: "Success: negative transaction",
			input: []byte(`
				{
					"account_id": 1,
					"operation_type_id": 1,
					"amount": 123.45
				}
			`),
			useCase: func() *mocks.MockUseCase {
				mockUseCase := mocks.NewMockUseCase(t)
				mockUseCase.On("Create", mock.Anything, mock.MatchedBy(func(txn domain.Transaction) bool {
					return txn.AccountID == 1 &&
						txn.OperationTypeID == 1 &&
						txn.Amount == 123.45 &&
						time.Since(txn.EventDate) < 2*time.Second
				})).Return(domain.Transaction{
					TransactionID:   1,
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          -123.45,
					EventDate:       eventDate,
				}, nil)
				return mockUseCase
			}(),
			output: output{
				statusCode: http.StatusCreated,
				messages: map[string]any{"data": domain.Transaction{
					TransactionID:   1,
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          -123.45,
					EventDate:       eventDate,
				}},
			},
		},
		{
			name: "Success: positive transaction",
			input: []byte(`
				{
					"account_id": 1,
					"operation_type_id": 4,
					"amount": 123.45
				}
			`),
			useCase: func() *mocks.MockUseCase {
				mockUseCase := mocks.NewMockUseCase(t)
				mockUseCase.On("Create", mock.Anything, mock.MatchedBy(func(txn domain.Transaction) bool {
					return txn.AccountID == 1 &&
						txn.OperationTypeID == 4 &&
						txn.Amount == 123.45 &&
						time.Since(txn.EventDate) < 2*time.Second
				})).Return(domain.Transaction{
					TransactionID:   1,
					AccountID:       1,
					OperationTypeID: 4,
					Amount:          123.45,
					EventDate:       eventDate,
				}, nil)
				return mockUseCase
			}(),
			output: output{
				statusCode: http.StatusCreated,
				messages: map[string]any{"data": domain.Transaction{
					TransactionID:   1,
					AccountID:       1,
					OperationTypeID: 4,
					Amount:          123.45,
					EventDate:       eventDate,
				}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, "/v1/transactions", bytes.NewReader(tt.input))
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			w := httptest.NewRecorder()

			e := echo.New()
			e.Validator = validator.NewValidator()

			c := e.NewContext(r, w)

			handler := NewHandler(tt.useCase)
			err := handler.Create(c)
			assert.NoError(t, err)

			assert.Equal(t, tt.output.statusCode, w.Code)

			messages, err := json.Marshal(tt.output.messages)
			assert.NoError(t, err)

			assert.Equal(t, string(messages), strings.TrimSpace(w.Body.String()))
		})
	}
}
