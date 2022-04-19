package service

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/hambyhacks/CrimsonIMS/app/models"
)

func Validate(u models.User) error {
	validate := validator.New()
	validate.RegisterValidation("username", NameValidation)
	validate.RegisterValidation("email", EmailValidation)
	return validate.Struct(u)
}

func NameValidation(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-zA-Z0-9_]+`)
	match := re.FindAllString(fl.Field().String(), -1)
	return len(match) == 1
}

func EmailValidation(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^(([a-zA-Z0-9_]{1,20})@([a-z]{1,20}).([a-z]{1,20}))$`)
	match := re.FindAllString(fl.Field().String(), -1)
	return len(match) == 1
}
