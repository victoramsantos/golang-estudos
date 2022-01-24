package repository

import (
	"sagapattern/waiter/domain"
	"sagapattern/waiter/repository/domain/queue"
)

type orderRepository struct {
	queue queue.Queue
}

func NewOrderRepository() domain.OrderRepository {
	return &orderRepository{
		queue: queue.NewQueue(),
	}
}

func (repository *orderRepository) MakeOrder(order *domain.Order) error {
	repository.queue.SendMessage(order)

	return nil
}
