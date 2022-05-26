package models

type Product struct {
	ID             int     `json:"id"`
	Name           string  `json:"product_name" validate:"product_name,min=4,max=70,required"`
	DeclaredPrice  float64 `json:"declared_price"`
	ShippingFee    float64 `json:"shipping_fee"`
	TrackingNumber string  `json:"tracking_number" validate:"tracking_number,required"`
	SellerName     string  `json:"seller_name"`
	SellerAddress  string  `json:"seller_address"`
	DateOrdered    string  `json:"date_ordered" validate:"datetime=2006-01-02"`
	DateReceived   string  `json:"date_received" validate:"datetime=2006-01-02"`
	ModeOfPayment  string  `json:"payment_mode"`
	StockCount     int     `json:"stock_count"`
}
