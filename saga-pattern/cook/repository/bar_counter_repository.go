package repository

import (
	"fmt"
	"sagapattern/cook/domain"
)

type barCounterRepository struct {
	// Kafka Consumer Topic
}

func NewBarCounterRepository() domain.BarCounterRepository {
	return &barCounterRepository{}
}

func (repository *barCounterRepository) Delivery(order *domain.Order) {
	fmt.Println("adding to bar counter ", order.OrderId)
}
