package request

type OrderRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
