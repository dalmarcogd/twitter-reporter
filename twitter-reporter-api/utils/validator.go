package utils

import "gopkg.in/go-playground/validator.v9"

type CustomValidator struct {
	Validator *validator.Validate
}

func NewCustomValidator(validator *validator.Validate) *CustomValidator {
	return &CustomValidator{Validator: validator}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
