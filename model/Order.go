package model

type Order struct {
	ID              int     `json:"id"`
	UsersId         string  `json:"users_id"`
	OrderDate       string  `json:"order_date"`
	Status          string  `json:"status"`
	TotalPrice      float64 `json:"total_price"`
	ShippingAddress string  `json:"shipping_address"`
	PaymentMethod   int     `json:"payment_method"`
}

type OrderItem struct {
	ID        int     `json:"id"`
	OrdersId  string  `json:"orders_id"`
	ProductId string  `json:"product_id"`
	Quantity  string  `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	Subtotal  string  `json:"subtotal"`
}
