package service

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/hambyhacks/CrimsonIMS/internal/data/models"
)

func Validate(p models.Product) error {
	validate := validator.New()
	validate.RegisterValidation("tracking_number", TrackingNumberValidation)
	validate.RegisterValidation("product_name", ProductNameValidation)
	return validate.Struct(p)
}

func TrackingNumberValidation(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^.([^a-z0-9]{0,5}[0-9]{12})$`)
	match := re.FindAllString(fl.Field().String(), -1)
	return len(match) == 1
}

func ProductNameValidation(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-zA-Z0-9\s]+`)
	match := re.FindAllString(fl.Field().String(), -1)
	return len(match) == 1
}
