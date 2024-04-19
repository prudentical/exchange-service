package util

import "github.com/go-playground/validator/v10"

type Validator interface {
	Validate(i interface{}) error
}

type PlaygroundValidator struct {
	validator *validator.Validate
}

func NewValidator() Validator {
	return &PlaygroundValidator{validator.New()}
}

func (v *PlaygroundValidator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return ValidationError{err.Error()}
	}
	return nil
}

type ValidationError struct {
	msg string
}

func (v ValidationError) Error() string {
	return v.msg
}
