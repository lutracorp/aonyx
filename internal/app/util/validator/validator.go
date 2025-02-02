package validator

import "github.com/go-playground/validator/v10"

// Validator validates data.
type Validator struct {
	validator *validator.Validate
}

// Validate validates structs.
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

// NewValidator creates new instance of Validator.
func NewValidator() *Validator {
	val := validator.New()

	_ = val.RegisterValidation("username", func(fl validator.FieldLevel) bool {
		return usernameRegex.MatchString(fl.Field().String())
	})

	return &Validator{val}
}
