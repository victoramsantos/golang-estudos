package usecase

import "sagapattern/waiter/domain"

type orderUsecase struct {
	repository domain.OrderRepository
}

func NewOrderUsecase(repository domain.OrderRepository) domain.OrderUsecase {
	return &orderUsecase{
		repository: repository,
	}
}

func (usecase *orderUsecase) OrderRequest(order *domain.Order) error {
	order.Status = "ORDERED"
	order.OrderId = usecase.repository.LastOrderNumber()
	return usecase.repository.MakeOrder(order)
}

func (usecase *orderUsecase) ListOrders() []domain.Order {
	return usecase.repository.ListOrders()
}
