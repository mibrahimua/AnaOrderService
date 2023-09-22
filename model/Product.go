package model

type Product struct {
	ID            int     `json:"id"`
	Category      string  `json:"category"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	Theme         string  `json:"theme"`
	StockQuantity int     `json:"stock_quantity"`
}
