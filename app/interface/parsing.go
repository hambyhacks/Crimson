package app

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hambyhacks/CrimsonIMS/app/interface/requests"
	"github.com/hambyhacks/CrimsonIMS/app/models"
)

func DecodeAddProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requests.AddProductRequest
	err := json.NewDecoder(r.Body).Decode(&req.Product)
	if err != nil {
		return nil, err
	}
	err = models.Validate(req.Product)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeGetAllProductsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requests.GetAllProductsRequest
	return req, nil
}

func DecodeGetProductByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requests.GetProductByIDRequest
	vars := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(vars, 10, 64)
	if err != nil {
		return nil, err
	}

	req = requests.GetProductByIDRequest{ID: int(id)}
	return req, nil
}

func DecodeDeleteProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requests.DeleteProductRequest
	vars := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(vars, 10, 64)
	if err != nil {
		return nil, err
	}

	req = requests.DeleteProductRequest{ID: int(id)}
	return req, nil
}

func DecodeUpdateProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requests.UpdateProductRequest
	err := json.NewDecoder(r.Body).Decode(&req.Product)
	if err != nil {
		return nil, err
	}
	err = models.Validate(req.Product)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func EncodeResponses(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
