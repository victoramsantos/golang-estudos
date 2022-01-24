package main

import (
	"sagapattern/cook/repository"
	"sagapattern/cook/usecase"
)

func main() {
	orderRepository := repository.NewOrderRepository()
	barCounterRepository := repository.NewBarCounterRepository()
	usecase := usecase.NewCookUsecase(
		barCounterRepository,
		orderRepository,
	)
	//Should be Cook() for ever
	usecase.Cook()
}
