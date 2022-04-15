package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/hambyhacks/CrimsonIMS/app/interface/products/requests"
	"github.com/hambyhacks/CrimsonIMS/app/interface/products/responses"
	prodsrv "github.com/hambyhacks/CrimsonIMS/service/products"
)

func MakeAddProductEndpoint(svc prodsrv.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(requests.AddProductRequest)
		msg, err := svc.AddProduct(ctx, req.Product)
		if err != nil {
			return responses.AddProductResponse{Msg: "unable to process request", Err: err}, err
		}
		return responses.AddProductResponse{Msg: msg, Err: nil}, nil
	}
}

func MakeGetProductByIDEndpoint(svc prodsrv.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(requests.GetProductByIDRequest)
		prodDetails, err := svc.GetProductByID(ctx, req.ID)
		if err != nil {
			return responses.GetProductByIDResponse{Product: nil, Err: err}, err
		}
		return responses.GetProductByIDResponse{Product: prodDetails, Err: nil}, nil
	}
}

func MakeGetAllProductsEndpoint(svc prodsrv.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		prodDetails, err := svc.GetAllProducts(ctx)
		if err != nil {
			return responses.GetAllProductsResponse{Product: nil, Err: err}, err
		}
		return responses.GetAllProductsResponse{Product: prodDetails, Err: nil}, nil
	}
}

func MakeDeleteProductEndpoint(svc prodsrv.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(requests.DeleteProductRequest)
		msg, err := svc.DeleteProduct(ctx, req.ID)
		if err != nil {
			return responses.DeleteProductResponse{Msg: "unable to process request", Err: err}, err
		}
		return responses.DeleteProductResponse{Msg: msg, Err: nil}, nil
	}
}

func MakeUpdateProductEndpoint(svc prodsrv.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(requests.UpdateProductRequest)
		msg, err := svc.UpdateProduct(ctx, req.Product)
		if err != nil {
			return responses.UpdateProductResponse{Msg: "unable to process request", Err: err}, err
		}
		return responses.UpdateProductResponse{Msg: msg, Err: nil}, nil
	}
}
