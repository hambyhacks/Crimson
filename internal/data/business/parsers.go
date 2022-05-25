package data

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
	ErrNotFound       = errors.New("product not found")
)

type errorer interface {
	error() error
}

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
func EncodeResponses(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	// Check errors in request before crafting response
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrValidation, ErrIntConv, ErrDBRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
