package model

type UnusedCart struct {
	ID        int    `json:"id"`
	UsersId   string `json:"users_id"`
	ProductId string `json:"product_id"`
	Quantity  string `json:"quantity"`
}
