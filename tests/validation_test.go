package tests

import (
	"testing"

	"github.com/hambyhacks/CrimsonIMS/internal/data/models"
	service "github.com/hambyhacks/CrimsonIMS/internal/service/products"
)

func TestValidation(t *testing.T) {
	p := models.Product{
		TrackingNumber: "JTXPH000000000000",
		Name:           "Test product",
	}

	err := service.Validate(p)
	if err != nil {
		t.Fatal(err)
	}
}
