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

type ProductStockDeducted struct {
	ProductStockId    int    `json:"product_stock_id"`
	ProductId         string `json:"product_id"`
	WarehouseId       string `json:"warehouse_id"`
	StockQuantity     string `json:"stock_quantity"`
	DeductedRemaining string `json:"deducted_remaining"`
	IsValid           bool   `json:"is_valid"`
}

type ProductStockCart struct {
	ID        int    `json:"id"`
	UsersId   string `json:"users_id"`
	ProductId string `json:"product_id"`
	Quantity  string `json:"quantity"`
}
