package usecase

import (
	"fmt"
	"sagapattern/cook/domain"
)

type cookUsecase struct {
	barCounterRepository domain.BarCounterRepository
	orderRepository      domain.OrderRepository
}

func NewCookUsecase(barCounterRepository domain.BarCounterRepository, orderRepository domain.OrderRepository) domain.CookUsecase {
	return &cookUsecase{
		barCounterRepository: barCounterRepository,
		orderRepository:      orderRepository,
	}
}

func (usecase *cookUsecase) Cook() {
	order := usecase.orderRepository.Consumes()
	//Should be a go func
	fmt.Println("cooking order", order)
	usecase.barCounterRepository.Delivery(&order)
}
