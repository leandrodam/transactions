package validator

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

func RegisterDecimalValidators(v *validator.Validate) {
	registerDecimalOperation("dgt", v, func(a, b *decimal.Decimal) bool {
		return a.GreaterThan(*b)
	})

	registerDecimalOperation("dgte", v, func(a, b *decimal.Decimal) bool {
		return a.GreaterThanOrEqual(*b)
	})

	registerDecimalOperation("dlt", v, func(a, b *decimal.Decimal) bool {
		return a.LessThan(*b)
	})

	registerDecimalOperation("dlte", v, func(a, b *decimal.Decimal) bool {
		return a.LessThanOrEqual(*b)
	})

	registerDecimalOperation("deq", v, func(a, b *decimal.Decimal) bool {
		return a.Equal(*b)
	})

	registerDecimalOperation("dneq", v, func(a, b *decimal.Decimal) bool {
		return !a.Equal(*b)
	})
	registerDecimalType(v)
}

func registerDecimalType(v *validator.Validate) {
	v.RegisterCustomTypeFunc(func(field reflect.Value) any {
		if valuer, ok := field.Interface().(decimal.Decimal); ok {
			return valuer.String()
		}
		return nil
	}, decimal.Decimal{})
}

func registerDecimalOperation(tag string, v *validator.Validate, comparator func(d1, d2 *decimal.Decimal) bool) {
	v.RegisterValidation(tag, func(fl validator.FieldLevel) bool {
		data, ok := fl.Field().Interface().(string)
		if !ok {
			return false
		}
		value, err := decimal.NewFromString(data)
		if err != nil {
			return false
		}
		baseValue, err := decimal.NewFromString(fl.Param())
		if err != nil {
			return false
		}
		return comparator(&value, &baseValue)
	})
}
