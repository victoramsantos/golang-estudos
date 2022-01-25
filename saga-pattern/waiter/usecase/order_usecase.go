package usecase

import (
	"sagapattern/waiter/domain"

	"github.com/google/uuid"
)

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
	order.OrderId = uuid.New().String()

	return usecase.repository.MakeOrder(order)
}
