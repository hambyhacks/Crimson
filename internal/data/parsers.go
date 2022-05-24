package app

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	prodreq "github.com/hambyhacks/CrimsonIMS/internal/data/interface/products/requests"
	prodValidator "github.com/hambyhacks/CrimsonIMS/internal/service/products"
)

// Error messages
var (
	ErrDecodingToJSON = errors.New("unable to decode to json")
	ErrEncodingToJSON = errors.New("unable to encode to json")
	ErrDBRequest      = errors.New("unable to process request")
	ErrValidation     = errors.New("validation failed")
	ErrIntConv        = errors.New("unable to convert string to int")
)

// Product Requests Decoder (line 30-85)
func DecodeAddProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req prodreq.AddProductRequest
	err := json.NewDecoder(r.Body).Decode(&req.Product)
	if err != nil {
		return nil, ErrDecodingToJSON
	}
	err = prodValidator.Validate(req.Product)
	if err != nil {
		return nil, ErrValidation
	}
	return req, nil
}

func DecodeGetAllProductsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req prodreq.GetAllProductsRequest
	return req, nil
}

func DecodeGetProductByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req prodreq.GetProductByIDRequest
	vars := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(vars, 10, 64)
	if err != nil {
		return nil, ErrIntConv
	}

	req = prodreq.GetProductByIDRequest{ID: int(id)}
	return req, nil
}

func DecodeDeleteProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req prodreq.DeleteProductRequest
	vars := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(vars, 10, 64)
	if err != nil {
		return nil, ErrIntConv
	}

	req = prodreq.DeleteProductRequest{ID: int(id)}
	return req, nil
}

func DecodeUpdateProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req prodreq.UpdateProductRequest
	err := json.NewDecoder(r.Body).Decode(&req.Product)
	if err != nil {
		return nil, ErrDecodingToJSON
	}
	err = prodValidator.Validate(req.Product)
	if err != nil {
		return nil, ErrValidation
	}
	return req, nil
}

// Response encoder
func EncodeResponses(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return ErrEncodingToJSON
	}
	return nil
}
