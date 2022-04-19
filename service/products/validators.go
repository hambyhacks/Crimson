package service

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/hambyhacks/CrimsonIMS/app/models"
)

func Validate(p models.Product) error {
	validate := validator.New()
	validate.RegisterValidation("sku", SKUValidation)
	validate.RegisterValidation("name", NameValidation)
	return validate.Struct(p)
}

func SKUValidation(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^([A-Z]{5}-[a-zA-Z0-9]{3}-[a-zA-Z0-9]{3}$)$`)
	match := re.FindAllString(fl.Field().String(), -1)
	return len(match) == 1
}

func NameValidation(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-zA-Z0-9\s]+`)
	match := re.FindAllString(fl.Field().String(), -1)
	return len(match) == 1
}
