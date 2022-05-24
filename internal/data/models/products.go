package models

import "time"

type Product struct {
	ID             int       `json:"id"`
	Name           string    `json:"product_name" validate:"name,min=4,max=70,required"`
	DeclaredPrice  float64   `json:"declared_price"`
	ShippingFee    float64   `json:"shipping_fee"`
	TrackingNumber string    `json:"tracking_number"`
	SellerName     string    `json:"seller_name"`
	SellerAddress  string    `json:"seller_address"`
	DateOrdered    time.Time `json:"-"`
	DateReceived   time.Time `json:"-"`
	ModeOfPayment  string    `json:"payment_mode"`
	StockCount     int       `json:"stock_count"`
}
