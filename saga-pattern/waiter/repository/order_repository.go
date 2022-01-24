package repository

import "sagapattern/waiter/domain"

type orderRepository struct {
	orders []domain.Order
}

func NewOrderRepository() domain.OrderRepository {
	return &orderRepository{
		orders: make([]domain.Order, 0),
	}
}

func (repository *orderRepository) MakeOrder(order *domain.Order) error {
	repository.orders = append(repository.orders, *order)

	return nil
}

func (repository *orderRepository) ListOrders() []domain.Order {
	return repository.orders
}

func (repository *orderRepository) LastOrderNumber() int {
	return len(repository.orders) + 1
}
