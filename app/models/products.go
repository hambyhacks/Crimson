package models

import "time"

type Product struct {
	ID           int       `json:"id"`
	Name         string    `json:"name" validate:"name,min=4,max=70,required"`
	Price        float64   `json:"price"`
	SKU          string    `json:"sku" validate:"sku,required"`
	DateOrdered  time.Time `json:"-"`
	DateReceived time.Time `json:"-"`
	StockCount   int       `json:"stock_count"`
}
