package validations

func RegistrationFormErrorMessage() map[string]string {
	return map[string]string{
		"required": "The field is required.",
		"email":    "The email must be valid",
		"min":      "should be more than the limit",
		"max":      "Should be less than the limit",
	}
}
