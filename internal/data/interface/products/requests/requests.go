package requests

import "github.com/hambyhacks/CrimsonIMS/internal/data/models"

type (
	AddProductRequest struct {
		Product models.Product
	}

	GetProductByIDRequest struct {
		ID int `json:"id"`
	}

	GetAllProductsRequest struct{}

	DeleteProductRequest struct {
		ID int `json:"id"`
	}

	UpdateProductRequest struct {
		Product models.Product
	}
)
