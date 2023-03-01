package validator

import (
	"gopkg.in/go-playground/validator.v9"
)

type _validator struct {
	validator *validator.Validate
}

// creates a new validator
func NewPlayGroundValidator() Validator {
	return &_validator{
		validator: validator.New(),
	}
}

// validates the json object passed as argument
func (v *_validator) ValidateJSON(data interface{}) (bool, error) {
	// Validates the JSON object and makes sure it meets the required request fields
	if err := v.validator.Struct(data); err != nil {
		return false, err
	}
	return true, nil
}
