package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/hambyhacks/CrimsonIMS/app/models"
)

// Error messages
var (
	RepoErr     = errors.New("unable to process database request")
	ErrNotFound = errors.New("product not found")
)

type ProductsRepository interface {
	AddProduct(ctx context.Context, products models.Products) error
	GetProductByID(ctx context.Context, id int) (interface{}, error)
	GetAllProducts(ctx context.Context) (interface{}, error)
	UpdateProduct(ctx context.Context, products models.Products) (string, error)
	DeleteProduct(ctx context.Context, id int) (string, error)
}

type prodRepo struct {
	db     *sql.DB
	logger log.Logger
}

func NewProdRepo(db *sql.DB, logger log.Logger) (ProductsRepository, error) {
	return &prodRepo{db: db, logger: log.With(logger, "repo", "postgresql")}, nil
}

// AddProduct implements ProductsRepository
func (r *prodRepo) AddProduct(ctx context.Context, products models.Products) error {
	q := `INSERT INTO products
		  (id, product_name, price, sku, date_ordered, date_received, stock_count)
		  VALUES ($1, $2, $3, $4, $5, $6, $7)`
	args := []interface{}{
		&products.ID,
		&products.Name,
		&products.Price,
		&products.SKU,
		&products.DateOrdered,
		&products.DateReceived,
		&products.StockCount,
	}

	_, err := r.db.ExecContext(ctx, q, args...)
	if err != nil {
		level.Error(r.logger).Log("repository-error", err)
		return err
	}
	return nil
}

// DeleteProduct implements ProductsRepository
func (r *prodRepo) DeleteProduct(ctx context.Context, id int) (string, error) {
	q := `DELETE FROM products where id = $1`
	res, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return "unable to process request", err
	}

	_, err = res.RowsAffected()
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return "", ErrNotFound
		default:
			return "", err
		}
	}
	return "successfully deleted product ", nil
}

// GetAllProducts implements ProductsRepository
func (r *prodRepo) GetAllProducts(ctx context.Context) (interface{}, error) {
	prod := models.Products{}
	var res []interface{}
	q := `SELECT 
		  id, product_name, price, sku, stock_count 
		  FROM products ORDER BY id DESC`

	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&prod.ID, &prod.Name, &prod.Price, &prod.SKU, &prod.StockCount)
		if err != nil {
			return nil, err
		}
		res = append([]interface{}{prod}, res...)
	}
	return res, nil
}

// GetProductByID implements ProductsRepository
func (r *prodRepo) GetProductByID(ctx context.Context, id int) (interface{}, error) {
	prod := models.Products{}
	q := `SELECT id, product_name, price, sku, stock_count FROM products WHERE id = $1`

	err := r.db.QueryRowContext(ctx, q, id).Scan(&prod.ID, &prod.Name, &prod.Price, &prod.SKU, &prod.StockCount)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}
	return prod, nil
}

// UpdateProduct implements ProductsRepository
func (r *prodRepo) UpdateProduct(ctx context.Context, products models.Products) (string, error) {
	q := `UPDATE products SET product_name = $1, price = $2, sku = $3, stock_count = $4 WHERE id = $5`
	args := []interface{}{products.Name, products.Price, products.SKU, products.StockCount, products.ID}

	res, err := r.db.ExecContext(ctx, q, args...)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return "", ErrNotFound
		default:
			return "", err
		}
	}

	_, err = res.RowsAffected()
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return "", ErrNotFound
		default:
			return "", err
		}
	}

	return "record successfully updated", nil

}
