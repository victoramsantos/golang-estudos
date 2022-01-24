package usecase

import (
	"fmt"
	"log"
	"sagapattern/cook/domain"
	"time"

	"github.com/spf13/viper"
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
	for {
		order, err := usecase.orderRepository.Consumes()
		if err != nil {
			log.Println("No message found, sleeping", viper.GetInt("app.sleep_period"), "seconds")
			time.Sleep(time.Duration(viper.GetInt("app.sleep_period")) * time.Second)
			continue
		}
		//Should be a go func
		fmt.Println("cooking order", order)
		usecase.barCounterRepository.Delivery(order)
	}
}
