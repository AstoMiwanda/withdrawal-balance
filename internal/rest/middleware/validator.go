package middleware

import (
	"errors"
	"fmt"
	"reflect"

	"gopkg.in/go-playground/validator.v9"
)

func IsRequestValid[T any](m *T) (bool, error) {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("json")
	})
	err := validate.Struct(m)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errMsg string
			for _, e := range validationErrors {
				switch e.Tag() {
				case "required":
					errMsg = fmt.Sprintf("%s is required", e.Field())
				case "gt":
					errMsg = fmt.Sprintf("%s must be greater than %s", e.Field(), e.Param())
				default:
					errMsg = fmt.Sprintf("%s is not valid", e.Field())
				}
				return false, errors.New(errMsg)
			}
		}
		return false, err
	}
	return true, nil
}
