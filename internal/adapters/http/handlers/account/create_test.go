package account

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	domain "github.com/leandrodam/transactions/internal/domain/account"
	"github.com/leandrodam/transactions/internal/infrastructure/exceptions"
	"github.com/leandrodam/transactions/internal/infrastructure/validator"
	mocks "github.com/leandrodam/transactions/internal/usecases/account/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Create(t *testing.T) {
	type output struct {
		statusCode int
		messages   map[string]any
	}

	tests := []struct {
		name    string
		input   []byte
		useCase *mocks.MockUseCase
		output  output
	}{
		{
			name:  "Error binding parameters",
			input: []byte(`{ "document_number": 12345678900 }`),
			output: output{
				statusCode: http.StatusBadRequest,
				messages:   exceptions.ErrBadRequest.ErrorJSON(),
			},
		},
		{
			name:  "Error: document number length",
			input: []byte(`{ "document_number": "123456789000" }`),
			output: output{
				statusCode: http.StatusBadRequest,
				messages:   map[string]any{"errors": []string{"DocumentNumber must be exactly 11 characters."}},
			},
		},
		{
			name:  "Error: document number is required",
			input: []byte(`{ "document_number": null }`),
			output: output{
				statusCode: http.StatusBadRequest,
				messages:   map[string]any{"errors": []string{"DocumentNumber is required."}},
			},
		},
		{
			name:  "Error: document number must be numeric",
			input: []byte(`{ "document_number": "1234567890A" }`),
			output: output{
				statusCode: http.StatusBadRequest,
				messages:   map[string]any{"errors": []string{"DocumentNumber must be numeric."}},
			},
		},
		{
			name:  "Error creating account",
			input: []byte(`{ "document_number": "12345678900" }`),
			useCase: func() *mocks.MockUseCase {
				mockUseCase := mocks.NewMockUseCase(t)
				mockUseCase.On("Create", mock.Anything, domain.Account{
					DocumentNumber: "12345678900",
				}).Return(domain.Account{}, exceptions.ErrInternal)
				return mockUseCase
			}(),
			output: output{
				statusCode: http.StatusInternalServerError,
				messages:   exceptions.ErrInternal.ErrorJSON(),
			},
		},
		{
			name:  "Success",
			input: []byte(`{ "document_number": "12345678900" }`),
			useCase: func() *mocks.MockUseCase {
				mockUseCase := mocks.NewMockUseCase(t)
				mockUseCase.On("Create", mock.Anything, domain.Account{
					DocumentNumber: "12345678900",
				}).Return(domain.Account{
					AccountID:      1,
					DocumentNumber: "12345678900",
				}, nil)
				return mockUseCase
			}(),
			output: output{
				statusCode: http.StatusCreated,
				messages: map[string]any{"data": domain.Account{
					AccountID:      1,
					DocumentNumber: "12345678900",
				}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, "/v1/accounts", bytes.NewReader(tt.input))
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
