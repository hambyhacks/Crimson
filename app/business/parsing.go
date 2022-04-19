package app

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	prodreq "github.com/hambyhacks/CrimsonIMS/app/interface/products/requests"
	userreq "github.com/hambyhacks/CrimsonIMS/app/interface/users/requests"
	"github.com/hambyhacks/CrimsonIMS/app/models"
	prodValidator "github.com/hambyhacks/CrimsonIMS/service/products"
	userValidator "github.com/hambyhacks/CrimsonIMS/service/users"
	"golang.org/x/crypto/bcrypt"
)

// Error Definitions
var (
	ErrDecodingToJSON  = errors.New("unable to decode to json")
	ErrEncodingToJSON  = errors.New("unable to encode to json")
	ErrDBRequest       = errors.New("unable to process request")
	ErrValidation      = errors.New("validation failed")
	ErrIntConv         = errors.New("unable to convert string to int")
	ErrEmailFieldEmpty = errors.New("empty email")
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

// Authentication Request Decoder (line 88-133)
func DecodeAddUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req userreq.AddUserRequest
	err := json.NewDecoder(r.Body).Decode(&req.User)
	if err != nil {
		return nil, ErrDecodingToJSON
	}
	err = userValidator.Validate(req.User)
	if err != nil {
		return nil, ErrValidation
	}

	passwd, err := SetPassword(req.User.Password.Plaintext)
	if err != nil {
		return nil, ErrDBRequest
	}

	req.User.Password.Hash = passwd
	req.User.Password.Plaintext = string(passwd)

	return req, nil
}

func DecodeGetUserByEmailRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req userreq.GetUserByEmailRequest
	email := r.URL.Query().Get("email")

	if req.Email == "" {
		return nil, ErrEmailFieldEmpty
	}

	req = userreq.GetUserByEmailRequest{Email: email}
	return req, nil
}

func DecodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req userreq.UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req.User)
	if err != nil {
		return nil, ErrDecodingToJSON
	}
	err = userValidator.Validate(req.User)
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

// Encryption functions
func SetPassword(plaintext string) ([]byte, error) {
	var user models.User
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintext), 12)
	if err != nil {
		return nil, err
	}
	user.Password.Plaintext = plaintext
	user.Password.Hash = hash
	return hash, nil
}

func CheckHash(plaintext string) (bool, error) {
	var user models.User
	err := bcrypt.CompareHashAndPassword(user.Password.Hash, []byte(plaintext))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}
