package services

import (
	"movie_premiuem/entity"
	"movie_premiuem/entity/repositories"
)

type OrderService interface {
	CreateOrder(order entity.Order) (entity.Order, error)
}

type orderService struct {
	orderRepository repositories.OrderRepository
}

func NewOrderService(orderRepository repositories.OrderRepository) OrderService {
	if orderRepository == nil {
		panic("orderRepository cannot be nil")
	}
	return &orderService{orderRepository: orderRepository}
}

func (o *orderService) CreateOrder(order entity.Order) (entity.Order, error) {
	// in future works this function can handle logic related to pre order
	// for example sms, cache , ...
	return o.orderRepository.CreateOrder(order)
}
