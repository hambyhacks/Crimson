package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	transport "github.com/go-kit/kit/transport/http"
	app "github.com/hambyhacks/CrimsonIMS/app/interface"
	endpoints "github.com/hambyhacks/CrimsonIMS/endpoints/products"
	prodsrv "github.com/hambyhacks/CrimsonIMS/service/products"
)

func NewHTTPHandler(svc prodsrv.ProductService) *chi.Mux {
	r := chi.NewRouter()

	// HTTP handlers
	AddProductHandler := transport.NewServer(
		endpoints.MakeAddProductEndpoint(svc),
		app.DecodeAddProductRequest,
		app.EncodeResponses,
	)

	GetAllProductsHandler := transport.NewServer(
		endpoints.MakeGetAllProductsEndpoint(svc),
		app.DecodeGetAllProductsRequest,
		app.EncodeResponses,
	)

	GetProductByIDHandler := transport.NewServer(
		endpoints.MakeGetProductByIDEndpoint(svc),
		app.DecodeGetProductByIDRequest,
		app.EncodeResponses,
	)

	DeleteProductHandler := transport.NewServer(
		endpoints.MakeDeleteProductEndpoint(svc),
		app.DecodeDeleteProductRequest,
		app.EncodeResponses,
	)

	UpdateProductHandler := transport.NewServer(
		endpoints.MakeUpdateProductEndpoint(svc),
		app.DecodeUpdateProductRequest,
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
				r.Method(http.MethodGet, "/products", GetAllProductsHandler)
				r.Method(http.MethodGet, "/products/{id:[0-9]+}", GetProductByIDHandler)
				r.Method(http.MethodDelete, "/products/delete/{id:[0-9]+}", DeleteProductHandler)
				r.Group(func(r chi.Router) {
					r.Method(http.MethodPost, "/products/add", AddProductHandler)
					r.Method(http.MethodPatch, "/products/update/{id:[0-9]+}", UpdateProductHandler)
				})
			})
		})
	})
	return r
}
