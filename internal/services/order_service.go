package services

import "github.com/idprm/go-ppob/internal/domain/repository"

type OrderService struct {
	orderRepo repository.IOrderRepository
}

func NewOrderService(orderRepo repository.IOrderRepository) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

type IOrderService interface {
}
