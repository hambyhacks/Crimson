package responses

type (
	AddProductResponse struct {
		Msg string `json:"msg"`
		Err error  `json:"error,omitempty"`
	}

	GetProductByIDResponse struct {
		Product interface{} `json:"product,omitempty"`
		Err     error       `json:"error,omitempty"`
	}

	GetAllProductsResponse struct {
		Product interface{} `json:"product,omitempty"`
		Err     error       `json:"error,omitempty"`
	}

	DeleteProductResponse struct {
		Msg string `json:"msg"`
		Err error  `json:"error,omitempty"`
	}

	UpdateProductResponse struct {
		Msg string `json:"msg"`
		Err error  `json:"error,omitempty"`
	}
)
