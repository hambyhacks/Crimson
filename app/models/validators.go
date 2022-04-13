package models

import (
	"regexp"

	"github.com/go-playground/validator"
)

func (p *Products) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", SKUValidation)
	return validate.Struct(p)
}

func SKUValidation(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^([A-Z]{5}-[a-zA-Z0-9]{3}-[a-zA-Z0-9]{3}$)$`)
	match := re.FindAllString(fl.Field().String(), -1)
	if len(match) != 1 {
		return false
	}
	return true
}
