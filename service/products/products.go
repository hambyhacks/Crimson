package service

import (
	"context"
	"log"
	"time"

	klog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/hambyhacks/CrimsonIMS/app/models"
)

type ProductService interface {
	AddProduct(ctx context.Context, products models.Products) (string, error)
	GetProductByID(ctx context.Context, id int) (interface{}, error)
	GetAllProducts(ctx context.Context) (interface{}, error)
	UpdateProduct(ctx context.Context, products models.Products) (string, error)
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
func (p *ProdServ) AddProduct(ctx context.Context, products models.Products) (string, error) {
	log.Println("[i] Endpoint: /v1/admin/add")
	logger := klog.With(p.logger, "method", "add product")
	msg := "successfully added product"
	prodDetails := models.Products{
		ID:           products.ID,
		Name:         products.Name,
		Price:        products.Price,
		SKU:          products.SKU,
		DateOrdered:  time.Now().UTC(),
		DateReceived: time.Now().UTC(),
		StockCount:   products.StockCount,
	}

	err := p.repo.AddProduct(ctx, prodDetails)
	if err != nil {
		level.Error(logger).Log("repository-error", err)
		return "unable to process request", err
	}
	return msg, nil
}

// DeleteProduct implements ProductService
func (p *ProdServ) DeleteProduct(ctx context.Context, id int) (string, error) {
	log.Println("[i] Endpoint: /v1/admin/delete/:id")
	logger := klog.With(p.logger, "method", "delete product")
	msg, err := p.repo.DeleteProduct(ctx, id)
	if err != nil {
		level.Error(logger).Log("repository-error", err)
		return "unable to process request", err
	}
	return msg, nil
}

// GetAllProducts implements ProductService
func (p *ProdServ) GetAllProducts(ctx context.Context) (interface{}, error) {
	log.Println("[i] Endpoint: /v1/admin/products")
	logger := klog.With(p.logger, "method", "get all products")
	var product interface{}

	product, err := p.repo.GetAllProducts(ctx)
	if err != nil {
		level.Error(logger).Log("repository-error", err)
		return nil, err
	}
	return product, nil
}

// GetProductByID implements ProductService
func (p *ProdServ) GetProductByID(ctx context.Context, id int) (interface{}, error) {
	log.Println("[i] Endpoint: /v1/admin/products/:id")
	logger := klog.With(p.logger, "method", "get product by id")
	var product interface{}

	product, err := p.repo.GetProductByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("repository-error", err)
		return nil, err
	}
	return product, nil
}

// UpdateProduct implements ProductService
func (p *ProdServ) UpdateProduct(ctx context.Context, products models.Products) (string, error) {
	log.Println("[i] Endpoint: /v1/admin/update/:id")
	logger := klog.With(p.logger, "method", "update product")
	msg := "successfully updated product details"

	prodDetails := models.Products{
		ID:           products.ID,
		Name:         products.Name,
		Price:        products.Price,
		SKU:          products.SKU,
		DateOrdered:  time.Now().UTC(),
		DateReceived: time.Now().UTC(),
		StockCount:   products.StockCount,
	}

	msg, err := p.repo.UpdateProduct(ctx, prodDetails)
	if err != nil {
		level.Error(logger).Log("repository-error", err)
		return "unable to process request", err
	}
	return msg, nil
}
