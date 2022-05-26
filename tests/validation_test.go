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
		DateOrdered:    "2022-05-30",
		DateReceived:   "2022-06-30",
	}

	err := service.Validate(p)
	if err != nil {
		t.Fatal(err)
	}
}
