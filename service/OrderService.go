package service

import (
	"AnaOrderService/model"
	"AnaOrderService/repository"
	"AnaOrderService/request"
)

type OrderService struct {
	productRepository *repository.OrderRepository
}

func NewOrderService(userRepository *repository.OrderRepository) *OrderService {
	return &OrderService{userRepository}
}

func (us *OrderService) CheckoutItems(param request.OrderRequest) ([]model.Product, error) {
	return us.productRepository.CheckoutItems(param)
}
