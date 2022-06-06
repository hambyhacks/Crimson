package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/log"

	"github.com/hambyhacks/CrimsonIMS/internal/data/models"
)

// Repo messages
var (
	ErrRepo           = errors.New("unable to process database request")
	ErrNotFound       = errors.New("product not found")
	ErrDuplicateEntry = errors.New("duplicate entry")
	ErrSeqReset       = errors.New("unable to reset sequence")
	RequestErr        = "unable to process database request"
	RequestSuccess    = "success"
)

type ProductsRepository interface {
	AddProduct(ctx context.Context, products models.Product) error
	GetProductByID(ctx context.Context, id int) (interface{}, error)
	GetAllProducts(ctx context.Context) (interface{}, error)
	UpdateProduct(ctx context.Context, products models.Product) (string, error)
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
func (r *prodRepo) AddProduct(ctx context.Context, products models.Product) error {
	q := `INSERT INTO products
		  (product_name, declared_price, shipping_fee, tracking_number, seller_name, seller_address, 
		  date_ordered, date_received, payment_mode, stock_count)
		  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
		  ON CONFLICT (product_name) DO NOTHING`

	args := []interface{}{
		&products.Name,
		&products.DeclaredPrice,
		&products.ShippingFee,
		&products.TrackingNumber,
		&products.SellerName,
		&products.SellerAddress,
		&products.DateOrdered,
		&products.DateReceived,
		&products.ModeOfPayment,
		&products.StockCount,
	}

	_, err := r.db.ExecContext(ctx, q, args...)
	if err != nil {
		return ErrRepo
	}

	_, err = r.db.ExecContext(ctx, "SELECT setval('products_id_seq',max(id)) FROM products")
	if err != nil {
		return ErrSeqReset
	}

	return nil
}

// DeleteProduct implements ProductsRepository
func (r *prodRepo) DeleteProduct(ctx context.Context, id int) (string, error) {
	q := `DELETE FROM products 
		  WHERE EXISTS (SELECT product_name FROM products WHERE product_name = $1)`
	res, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return RequestErr, ErrRepo
	}

	sqlres, err := res.RowsAffected()
	if sqlres == 0 {
		return RequestErr, ErrNotFound
	}

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return RequestErr, ErrNotFound
		default:
			return RequestErr, ErrRepo
		}
	}

	_, err = r.db.ExecContext(ctx, "SELECT setval('products_id_seq',max(id)) FROM products")
	if err != nil {
		return RequestErr, ErrSeqReset
	}

	return RequestSuccess, nil
}

// GetAllProducts implements ProductsRepository
func (r *prodRepo) GetAllProducts(ctx context.Context) (interface{}, error) {
	prod := models.Product{}
	var res []interface{}
	q := `SELECT 
		  id,product_name, declared_price, shipping_fee, 
		  tracking_number, seller_name,
		  seller_address, date_ordered, date_received,
		  payment_mode, stock_count 
		  FROM products ORDER BY id DESC`

	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return RequestErr, ErrNotFound
		default:
			return RequestErr, ErrRepo
		}
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&prod.ID, &prod.Name, &prod.DeclaredPrice, &prod.ShippingFee, &prod.TrackingNumber, &prod.SellerName, &prod.SellerAddress, &prod.DateOrdered, &prod.DateReceived, &prod.ModeOfPayment, &prod.StockCount)
		if err != nil {
			return RequestErr, ErrRepo
		}
		res = append([]interface{}{prod}, res...)
	}
	return res, nil
}

// GetProductByID implements ProductsRepository
func (r *prodRepo) GetProductByID(ctx context.Context, id int) (interface{}, error) {
	prod := models.Product{}
	q := `SELECT id, product_name, declared_price, 
		  shipping_fee, tracking_number, seller_name,
		  seller_address, date_ordered, date_received,
		  payment_mode, stock_count FROM products WHERE id = $1`

	err := r.db.QueryRowContext(ctx, q, id).Scan(&prod.ID, &prod.Name, &prod.DeclaredPrice, &prod.ShippingFee, &prod.TrackingNumber, &prod.SellerName, &prod.SellerAddress, &prod.DateOrdered, &prod.DateReceived, &prod.ModeOfPayment, &prod.StockCount)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return RequestErr, ErrNotFound
		default:
			return RequestErr, ErrRepo
		}
	}
	return prod, nil
}

// UpdateProduct implements ProductsRepository
func (r *prodRepo) UpdateProduct(ctx context.Context, products models.Product) (string, error) {
	q := `UPDATE products SET 
		  product_name = $1, declared_price = $2, 
		  shipping_fee = $3, tracking_number = $4, 
		  seller_name = $5, seller_address = $6,
		  date_ordered = $7, date_received = $8,
		  payment_mode = $9, stock_count = $10 WHERE id = $11`
	args := []interface{}{
		products.Name,
		products.DeclaredPrice,
		products.ShippingFee,
		products.TrackingNumber,
		products.SellerName,
		products.SellerAddress,
		products.DateOrdered,
		products.DateReceived,
		products.ModeOfPayment,
		products.StockCount,
		products.ID,
	}

	res, err := r.db.ExecContext(ctx, q, args...)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return RequestErr, ErrNotFound
		default:
			return RequestErr, ErrRepo
		}
	}

	_, err = res.RowsAffected()
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return RequestErr, ErrNotFound
		default:
			return RequestErr, ErrRepo
		}
	}
	return RequestSuccess, nil
}
