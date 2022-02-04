package helper

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field   interface{} `json:"field"`
	Message string      `json:"message"`
}

func customMessageErrorForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "gte":
		return fmt.Sprintf("Length must be greater than %s", fe.Param())
	}
	return fe.Error() // default error
}

var v *validator.Validate = validator.New()

func Validate(object interface{}) (bool, []ApiError) {
	var ve validator.ValidationErrors
	err := v.Struct(object)

	if err != nil && errors.As(err, &ve) {
		out := make([]ApiError, len(ve))
		for i, fe := range ve {
			out[i] = ApiError{fe.Field(), customMessageErrorForTag(fe)}
		}
		return false, out
	}

	return true, nil
}

func NewSingleError(message string) []ApiError {
	out := make([]ApiError, 1)
	out[0] = ApiError{nil, message}
	return out
}
