package helper

import (
	"github.com/go-playground/validator/v10"
)

var v *validator.Validate = validator.New()

func Validate(object interface{}) (bool, string) {
	err := v.Struct(object)
	if err != nil {
		return false, err.Error()
	}
	return true, ""
}
