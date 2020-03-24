package utils

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
	"testing"
)

func TestCustomValidator_Validate(t *testing.T) {
	type ReportersRequest struct {
		Tag string `json:"tag" validate:"required"`
	}
	a := ReportersRequest{}
	assert.Error(t, NewCustomValidator(validator.New()).Validate(a))
	a.Tag = "12345678900"
	assert.NoError(t, NewCustomValidator(validator.New()).Validate(a))
}
