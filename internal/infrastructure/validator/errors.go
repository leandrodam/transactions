package validator

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	validator     *validator.Validate
	translator    ut.Translator
	originalError error
}

func (ve *ValidationError) Error() string {
	return ve.originalError.Error()
}

func (ve *ValidationError) GetMessages() []string {
	var messages []string

	if validationErrors, ok := ve.originalError.(validator.ValidationErrors); ok {
		for _, err := range validationErrors {
			messages = append(messages, err.Translate(ve.translator))
		}
	}

	return messages
}
