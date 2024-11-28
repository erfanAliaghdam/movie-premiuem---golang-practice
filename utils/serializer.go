package utils

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func FieldValidator(err error) map[string]string {
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
