package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	transport "github.com/go-kit/kit/transport/http"
	app "github.com/hambyhacks/CrimsonIMS/app/business"
	prodEndpoints "github.com/hambyhacks/CrimsonIMS/endpoints/products"
	authEndpoints "github.com/hambyhacks/CrimsonIMS/endpoints/users"
	prodsrv "github.com/hambyhacks/CrimsonIMS/service/products"
	usersrv "github.com/hambyhacks/CrimsonIMS/service/users"
)

func NewHTTPHandler(prodsvc prodsrv.ProductService, authsvc usersrv.UserService) *chi.Mux {
	r := chi.NewRouter()

	// HTTP handlers
	AddProductHandler := transport.NewServer(
		prodEndpoints.MakeAddProductEndpoint(prodsvc),
		app.DecodeAddProductRequest,
		app.EncodeResponses,
	)

	GetAllProductsHandler := transport.NewServer(
		prodEndpoints.MakeGetAllProductsEndpoint(prodsvc),
		app.DecodeGetAllProductsRequest,
		app.EncodeResponses,
	)

	GetProductByIDHandler := transport.NewServer(
		prodEndpoints.MakeGetProductByIDEndpoint(prodsvc),
		app.DecodeGetProductByIDRequest,
		app.EncodeResponses,
	)

	DeleteProductHandler := transport.NewServer(
		prodEndpoints.MakeDeleteProductEndpoint(prodsvc),
		app.DecodeDeleteProductRequest,
		app.EncodeResponses,
	)

	UpdateProductHandler := transport.NewServer(
		prodEndpoints.MakeUpdateProductEndpoint(prodsvc),
		app.DecodeUpdateProductRequest,
		app.EncodeResponses,
	)

	AddUserHandler := transport.NewServer(
		authEndpoints.MakeAddUserEndpoint(authsvc),
		app.DecodeAddUserRequest,
		app.EncodeResponses,
	)

	GetUserByEmailHandler := transport.NewServer(
		authEndpoints.MakeGetUserByEmailEndpoint(authsvc),
		app.DecodeGetUserByEmailRequest,
		app.EncodeResponses,
	)

	UpdateUserHandler := transport.NewServer(
		authEndpoints.MakeUpdateUserEndpoint(authsvc),
		app.DecodeUpdateUserRequest,
		app.EncodeResponses,
	)

	// Public routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/v1/info", http.StatusMovedPermanently)
	})

	r.Get("/v1", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/v1/info", http.StatusMovedPermanently)
	})

	r.Route("/v1", func(r chi.Router) {
		r.Get("/info", Index)

		// Private routes
		r.Group(func(r chi.Router) {
			// Add Authentication middleware here
			r.Route("/admin", func(r chi.Router) {
				// Products Service
				r.Group(func(r chi.Router) {
					r.Method(http.MethodGet, "/products", GetAllProductsHandler)
					r.Method(http.MethodGet, "/products/{id:[0-9]+}", GetProductByIDHandler)
					r.Method(http.MethodDelete, "/products/delete/{id:[0-9]+}", DeleteProductHandler)
					r.Method(http.MethodPost, "/products/add", AddProductHandler)
					r.Method(http.MethodPatch, "/products/update/{id:[0-9]+}", UpdateProductHandler)
				})
				// Auth service
				r.Group(func(r chi.Router) {
					r.Method(http.MethodGet, "/users/", GetUserByEmailHandler) // unfinished: needs to get query parameter
					r.Method(http.MethodPost, "/users/add", AddUserHandler)
					r.Method(http.MethodPatch, "/users/update", UpdateUserHandler) // unused
				})
			})
		})
	})
	return r
}
