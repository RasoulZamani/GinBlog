package errors

import (
	"errors"

	"github.com/RasoulZamani/hiGin/internal/providers/validations"
	"github.com/go-playground/validator/v10"
)

var errorsList map[string]string

func Init() {
	errorsList = map[string]string{}
}

func SetFromErrors(err error) {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			errorsList[fieldError.Field()] = GetErrorMessage(fieldError.Tag())
		}
	}
}

func Add(key string, value string) {
	errorsList[key] = value
}

// this function get tag and get message for it; e.g.:"required" -> this field is required
func GetErrorMessage(tag string) string {
	return validations.RegistrationFormErrorMessage()[tag]

}

func Get() map[string]string {
	return errorsList
}
