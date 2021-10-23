package validation

import validation "github.com/go-ozzo/ozzo-validation"

func ValidateUserCreate(data model.UserRequest) error {
	return validation.Errors{
		"name": validation.Validate(data.Name, validation.Required),
	}.Filter()
}
