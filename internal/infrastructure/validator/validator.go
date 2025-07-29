package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator  *validator.Validate
	translator ut.Translator
}

func NewValidator() echo.Validator {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	validate := validator.New()

	setupCustomMessages(validate, trans)

	return &CustomValidator{
		validator:  validate,
		translator: trans,
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return &ValidationError{
			validator:     cv.validator,
			translator:    cv.translator,
			originalError: err,
		}
	}

	return nil
}

func (cv *CustomValidator) GetTranslator() ut.Translator {
	return cv.translator
}

func setupCustomMessages(validate *validator.Validate, trans ut.Translator) {
	messages := map[string]string{
		"required": "{0} is required.",
		"numeric":  "{0} must be numeric.",
		"len":      "{0} must be exactly {1} characters.",
		"min":      "{0} must be at least {1} characters.",
		"max":      "{0} cannot exceed {1} characters.",
		"gt":       "{0} must be greater than {1}.",
		"gte":      "{0} must be greater than or equal to {1}.",
	}

	for tag, msg := range messages {
		registerMessage(validate, trans, tag, msg)
	}
}

func registerMessage(validate *validator.Validate, trans ut.Translator, tag, message string) {
	validate.RegisterTranslation(
		tag,
		trans,
		func(ut ut.Translator) error {
			return ut.Add(tag, message, true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(tag, fe.Field(), fe.Param())
			return t
		},
	)
}
