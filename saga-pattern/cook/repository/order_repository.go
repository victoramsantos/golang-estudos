package repository

import (
	"sagapattern/cook/domain"
	"sagapattern/cook/repository/domain/queue"
)

type orderRepository struct {
	queue queue.Queue
}

func NewOrderRepository() domain.OrderRepository {
	return &orderRepository{
		queue: queue.NewQueue(),
	}
}

func (repository *orderRepository) Consumes() domain.Order {
	return repository.queue.ReadMessage()
}
