package validation_service

import (
	"github.com/go-playground/validator/v10"
	"github.com/rahmanbesir/EventCollector/internal_error"
)

type Service interface {
	Validate(s interface{}) error
}

type validationService struct {
	validate *validator.Validate
}

func New(validate *validator.Validate) Service {
	return &validationService{
		validate: validate,
	}
}

func (v *validationService) Validate(s interface{}) error {
	err := v.validate.Struct(s)

	if err != nil {
		return internal_error.CreateValidationError(err)
	}
	return nil
}
