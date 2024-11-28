package service

import (
	"movie_premiuem/order/entity"
	"movie_premiuem/order/repository"
)

type OrderService interface {
	CreateOrder(order entity.Order) (entity.Order, error)
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) OrderService {
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
