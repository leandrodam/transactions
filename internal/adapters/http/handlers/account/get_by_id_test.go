package account

import (
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

func Test_GetById(t *testing.T) {
	type output struct {
		statusCode int
		messages   map[string]any
	}

	tests := []struct {
		name    string
		input   string
		useCase *mocks.MockUseCase
		output  output
	}{
		{
			name:  "Error parsing accountID",
			input: "invalid",
			output: output{
				statusCode: http.StatusBadRequest,
				messages:   exceptions.ErrInvalidAccountID.ErrorJSON(),
			},
		},
		{
			name:  "Error fetching account",
			input: "1",
			useCase: func() *mocks.MockUseCase {
				mockUseCase := mocks.NewMockUseCase(t)
				mockUseCase.On("GetByID", mock.Anything, 1).
					Return(domain.Account{}, exceptions.ErrInternal)
				return mockUseCase
			}(),
			output: output{
				statusCode: http.StatusInternalServerError,
				messages:   exceptions.ErrInternal.ErrorJSON(),
			},
		},
		{
			name:  "Success",
			input: "1",
			useCase: func() *mocks.MockUseCase {
				mockUseCase := mocks.NewMockUseCase(t)
				mockUseCase.On("GetByID", mock.Anything, 1).
					Return(domain.Account{
						AccountID:      1,
						DocumentNumber: "12345678900",
					}, nil)
				return mockUseCase
			}(),
			output: output{
				statusCode: http.StatusOK,
				messages: map[string]any{"data": domain.Account{
					AccountID:      1,
					DocumentNumber: "12345678900",
				}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, "/v1/accounts/:accountId", nil)
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			w := httptest.NewRecorder()

			e := echo.New()
			e.Validator = validator.NewValidator()

			c := e.NewContext(r, w)
			c.SetParamNames("accountId")
			c.SetParamValues(tt.input)

			handler := NewHandler(tt.useCase)
			err := handler.GetByID(c)
			assert.NoError(t, err)

			assert.Equal(t, tt.output.statusCode, w.Code)

			messages, err := json.Marshal(tt.output.messages)
			assert.NoError(t, err)

			assert.Equal(t, string(messages), strings.TrimSpace(w.Body.String()))
		})
	}
}
