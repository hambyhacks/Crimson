package tests

import (
	"testing"

	"github.com/hambyhacks/CrimsonIMS/app/models"
	service "github.com/hambyhacks/CrimsonIMS/service/products"
)

func TestSKUValidation(t *testing.T) {
	p := models.Products{
		ID:         1,
		Name:       "test",
		Price:      30.0,
		SKU:        "AMDRX-580-001",
		StockCount: 3,
	}

	err := service.Validate(p)
	if err != nil {
		t.Fatal(err)
	}
}
