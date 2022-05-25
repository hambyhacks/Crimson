package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	transport "github.com/go-kit/kit/transport/http"
	parsers "github.com/hambyhacks/CrimsonIMS/internal/data/business"
	prodEndpoints "github.com/hambyhacks/CrimsonIMS/internal/endpoints/products"
	prodsrv "github.com/hambyhacks/CrimsonIMS/internal/service/products"
)

func NewHTTPHandler(prodsvc prodsrv.ProductService) *chi.Mux {
	r := chi.NewRouter()

	// HTTP handlers
	AddProductHandler := transport.NewServer(
		prodEndpoints.MakeAddProductEndpoint(prodsvc),
		parsers.DecodeAddProductRequest,
		parsers.EncodeResponses,
	)

	GetAllProductsHandler := transport.NewServer(
		prodEndpoints.MakeGetAllProductsEndpoint(prodsvc),
		parsers.DecodeGetAllProductsRequest,
		parsers.EncodeResponses,
	)

	GetProductByIDHandler := transport.NewServer(
		prodEndpoints.MakeGetProductByIDEndpoint(prodsvc),
		parsers.DecodeGetProductByIDRequest,
		parsers.EncodeResponses,
	)

	DeleteProductHandler := transport.NewServer(
		prodEndpoints.MakeDeleteProductEndpoint(prodsvc),
		parsers.DecodeDeleteProductRequest,
		parsers.EncodeResponses,
	)

	UpdateProductHandler := transport.NewServer(
		prodEndpoints.MakeUpdateProductEndpoint(prodsvc),
		parsers.DecodeUpdateProductRequest,
		parsers.EncodeResponses,
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
			})
		})
	})
	return r
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Crimson IMS\n"))
}
