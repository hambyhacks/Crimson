package app

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	authreq "github.com/hambyhacks/CrimsonIMS/app/interface/auth/requests"
	prodreq "github.com/hambyhacks/CrimsonIMS/app/interface/products/requests"
	"github.com/hambyhacks/CrimsonIMS/app/models"
	authValidator "github.com/hambyhacks/CrimsonIMS/service/auth"
	prodValidator "github.com/hambyhacks/CrimsonIMS/service/products"
	"golang.org/x/crypto/bcrypt"
)

func DecodeAddProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req prodreq.AddProductRequest
	err := json.NewDecoder(r.Body).Decode(&req.Product)
	if err != nil {
		return nil, err
	}
	err = prodValidator.Validate(req.Product)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	req = prodreq.GetProductByIDRequest{ID: int(id)}
	return req, nil
}

func DecodeDeleteProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req prodreq.DeleteProductRequest
	vars := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(vars, 10, 64)
	if err != nil {
		return nil, err
	}

	req = prodreq.DeleteProductRequest{ID: int(id)}
	return req, nil
}

func DecodeUpdateProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req prodreq.UpdateProductRequest
	err := json.NewDecoder(r.Body).Decode(&req.Product)
	if err != nil {
		return nil, err
	}
	err = prodValidator.Validate(req.Product)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeAddUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req authreq.AddUserRequest
	err := json.NewDecoder(r.Body).Decode(&req.User)
	if err != nil {
		return nil, err
	}
	err = authValidator.Validate(req.User)
	if err != nil {
		return nil, err
	}

	hash, err := SetPassword(req.User.Password.Plaintext)
	if err != nil {
		return nil, err
	}

	req.User.Password.Plaintext = hash

	return req, nil
}

func DecodeGetUserByEmailRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req authreq.GetUserByEmailRequest
	email := r.URL.Query().Get("email")

	if req.Email == "" {
		return nil, errors.New("empty email")
	}

	req = authreq.GetUserByEmailRequest{Email: email}
	return req, nil
}

func DecodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req authreq.UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req.User)
	if err != nil {
		return nil, err
	}
	err = authValidator.Validate(req.User)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func EncodeResponses(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func SetPassword(plaintext string) (string, error) {
	var user models.User
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintext), 12)
	if err != nil {
		return "", err
	}
	user.Password.Plaintext = plaintext
	user.Password.Hash = hash
	return string(hash), nil
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
