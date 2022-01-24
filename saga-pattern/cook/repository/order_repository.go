package repository

import (
	"fmt"
	"sagapattern/cook/domain"
)

type orderRepository struct {
	// Kafka Consumer Topic
}

func NewOrderRepository() domain.OrderRepository {
	return &orderRepository{}
}

func (repository *orderRepository) Consumes() domain.Order {
	fmt.Println("Consuming order", 1)
	items := make([]int, 0)
	items = append(items, 1)
	items = append(items, 5)

	return domain.Order{
		OrderId: 1,
		Items:   items,
	}
}
