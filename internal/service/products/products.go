package service

import (
	"context"
	"fmt"

	klog "github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"github.com/hambyhacks/CrimsonIMS/internal/data/models"
)

type ProductService interface {
	AddProduct(ctx context.Context, products models.Product) (string, error)
	GetProductByID(ctx context.Context, id int) (interface{}, error)
	GetAllProducts(ctx context.Context) (interface{}, error)
	UpdateProduct(ctx context.Context, products models.Product) (string, error)
	DeleteProduct(ctx context.Context, id int) (string, error)
}

// Implementation of Product Service interface
type ProdServ struct {
	repo   ProductsRepository
	logger klog.Logger
}

func NewProdServ(repo ProductsRepository, logger klog.Logger) ProductService {
	return &ProdServ{
		repo:   repo,
		logger: logger,
	}
}

// AddProduct implements ProductService
func (p *ProdServ) AddProduct(ctx context.Context, products models.Product) (string, error) {
	logger := klog.With(p.logger, "method", "add product")

	level.Debug(logger).Log("endpoint", "/v1/admin/products/add")
	prodDetails := models.Product{
		ID:             products.ID,
		Name:           products.Name,
		DeclaredPrice:  products.DeclaredPrice,
		ShippingFee:    products.ShippingFee,
		TrackingNumber: products.TrackingNumber,
		SellerName:     products.SellerName,
		SellerAddress:  products.SellerAddress,
		DateOrdered:    products.DateOrdered,
		DateReceived:   products.DateReceived,
		ModeOfPayment:  products.ModeOfPayment,
		StockCount:     products.StockCount,
	}

	err := p.repo.AddProduct(ctx, prodDetails)
	if err != nil {
		level.Error(logger).Log("repository-err", err)
		return RequestErr, err
	}
	return RequestSuccess, nil
}

// DeleteProduct implements ProductService
func (p *ProdServ) DeleteProduct(ctx context.Context, id int) (string, error) {
	logger := klog.With(p.logger, "method", "delete product")

	level.Debug(logger).Log("endpoint", "/v1/admin/products/delete")
	msg, err := p.repo.DeleteProduct(ctx, id)
	if err != nil {
		level.Error(logger).Log("repository-error", err)
		return RequestErr, err
	}
	return msg, nil
}

// GetAllProducts implements ProductService
func (p *ProdServ) GetAllProducts(ctx context.Context) (interface{}, error) {
	var product interface{}
	logger := klog.With(p.logger, "method", "get all products")

	level.Debug(logger).Log("endpoint", "/v1/admin/products")
	product, err := p.repo.GetAllProducts(ctx)
	if err != nil {
		level.Error(logger).Log("repository-error", err)
		return RequestErr, err
	}
	return product, nil
}

// GetProductByID implements ProductService
func (p *ProdServ) GetProductByID(ctx context.Context, id int) (interface{}, error) {
	logger := klog.With(p.logger, "method", "get product by id")
	level.Debug(logger).Log("endpoint", "/v1/admin/products"+fmt.Sprintf("/%d", id))

	product, err := p.repo.GetProductByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("repository-error", err)
		return RequestErr, err
	}
	return product, nil
}

// UpdateProduct implements ProductService
func (p *ProdServ) UpdateProduct(ctx context.Context, products models.Product) (string, error) {
	logger := klog.With(p.logger, "method", "update product")

	level.Debug(logger).Log("endpoint", "/v1/admin/products/update")
	prodDetails := models.Product{
		ID:             products.ID,
		Name:           products.Name,
		DeclaredPrice:  products.DeclaredPrice,
		ShippingFee:    products.ShippingFee,
		TrackingNumber: products.TrackingNumber,
		SellerName:     products.SellerName,
		SellerAddress:  products.SellerAddress,
		DateOrdered:    products.DateOrdered,
		DateReceived:   products.DateReceived,
		ModeOfPayment:  products.ModeOfPayment,
		StockCount:     products.StockCount,
	}

	msg, err := p.repo.UpdateProduct(ctx, prodDetails)
	if err != nil {
		level.Error(logger).Log("repository-error", err)
		return RequestErr, err
	}
	return msg, nil
}
