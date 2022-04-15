package tests

import (
	"testing"

	"github.com/hambyhacks/CrimsonIMS/app/models"
	service "github.com/hambyhacks/CrimsonIMS/service"
)

func TestValidation(t *testing.T) {
	p := models.Products{
		ID:         1,
		Name:       "AMD Radeon RX570",
		Price:      3000.0,
		SKU:        "AMDRX-570-001",
		StockCount: 3,
	}

	err := service.Validate(p)
	if err != nil {
		t.Fatal(err)
	}
}
