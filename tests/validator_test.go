package tests

import (
	"testing"

	"github.com/hambyhacks/CrimsonIMS/app/models"
	prodValidator "github.com/hambyhacks/CrimsonIMS/service/products"
)

func TestValidation(t *testing.T) {
	p := models.Product{
		ID:         1,
		Name:       "AMD Radeon RX570",
		Price:      3000.0,
		SKU:        "AMDRX-570-001",
		StockCount: 3,
	}

	err := prodValidator.Validate(p)
	if err != nil {
		t.Fatal(err)
	}
}
