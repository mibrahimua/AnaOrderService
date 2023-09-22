package request

type OrderRequest struct {
	UserId    int `json:"users_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
