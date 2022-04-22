package tests

import (
	"testing"

	"github.com/hambyhacks/CrimsonIMS/app/models"
	prodValidator "github.com/hambyhacks/CrimsonIMS/service/products"
	userValidator "github.com/hambyhacks/CrimsonIMS/service/users"
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

	u := models.User{
		ID:    1,
		Name:  "test",
		Email: "test@gmail.com",
		Password: models.Password{
			Plaintext: "$2a$12$WTtOpCIP26eAzR5b4shCteQXkRoZuZLJJn1W3lKZJlYK9yp0R1BXm",
			Hash:      []byte("$2a$12$WTtOpCIP26eAzR5b4shCteQXkRoZuZLJJn1W3lKZJlYK9yp0R1BXm"),
		},
		Activated: false,
	}

	err = userValidator.Validate(u)
	if err != nil {
		t.Fatal(err)
	}
}
