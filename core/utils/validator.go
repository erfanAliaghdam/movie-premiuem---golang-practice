package utils

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func fieldValidator(err error) map[string]string {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		fieldErrors := make(map[string]string)
		for _, fieldErr := range validationErrors {
			fieldErrors[fieldErr.Field()] = fmt.Sprintf(fieldErr.Tag())
		}
		return fieldErrors
	}

	return nil
}

func ValidateField(structForValidation interface{}) (bool, map[string]string) {
	validate := validator.New()
	err := validate.Struct(structForValidation)
	if err != nil {
		fields := fieldValidator(err)
		if fields != nil {
			return false, fields
		}
		return false, nil
	}
	return true, map[string]string{}

}
