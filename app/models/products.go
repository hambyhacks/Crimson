package models

import "time"

type Products struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	SKU          string    `json:"sku"`
	DateOrdered  time.Time `json:"-"`
	DateReceived time.Time `json:"-"`
	StockCount   int       `json:"stock"`
}
