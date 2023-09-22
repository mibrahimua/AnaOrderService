package service

import (
	"AnaOrderService/model"
	"AnaOrderService/repository"
)

type OrderService struct {
	productRepository *repository.ProductRepository
}

func NewOrderService(userRepository *repository.ProductRepository) *OrderService {
	return &OrderService{userRepository}
}

func (us *OrderService) CheckoutItems(productName string) ([]model.Product, error) {
	return us.productRepository.CheckoutItems(productName)
}

/**
user send id product and quantity
order service processed it to deduction its stocks and insert into table orders and order_items with db transaction
return success if deduction are valid with response detail order
return false if deduction are invalid because stocks are outnumber
*/
