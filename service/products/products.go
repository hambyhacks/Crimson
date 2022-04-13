package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
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
type prodServ struct {
	repo   ProductsRepository
	logger log.Logger
}

func NewProdServ(repo ProductsRepository, logger log.Logger) ProductService {
	return &prodServ{
		repo:   repo,
		logger: logger,
	}
}

// AddProduct implements ProductService
func (p *prodServ) AddProduct(ctx context.Context, products models.Products) (string, error) {
	logger := log.With(p.logger, "method", "add product")
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
func (p *prodServ) DeleteProduct(ctx context.Context, id int) (string, error) {
	logger := log.With(p.logger, "method", "delete product")
	msg, err := p.repo.DeleteProduct(ctx, id)
	if err != nil {
		level.Error(logger).Log("repository-error", err)
		return "unable to process request", err
	}
	return msg, nil
}

// GetAllProducts implements ProductService
func (p *prodServ) GetAllProducts(ctx context.Context) (interface{}, error) {
	logger := log.With(p.logger, "method", "get all products")
	var product interface{}

	product, err := p.GetAllProducts(ctx)
	if err != nil {
		level.Error(logger).Log("repository-error", err)
		return nil, err
	}
	return product, nil
}

// GetProductByID implements ProductService
func (p *prodServ) GetProductByID(ctx context.Context, id int) (interface{}, error) {
	logger := log.With(p.logger, "method", "get product by id")
	var product interface{}

	product, err := p.repo.GetProductByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("repository-error", err)
		return nil, err
	}
	return product, nil
}

// UpdateProduct implements ProductService
func (p *prodServ) UpdateProduct(ctx context.Context, products models.Products) (string, error) {
	logger := log.With(p.logger, "method", "update product")
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
