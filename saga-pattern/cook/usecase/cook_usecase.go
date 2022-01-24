package usecase

import (
	"log"
	"math/rand"
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
		cooking(order)
		usecase.barCounterRepository.Delivery(order)
	}
}

func cooking(order *domain.Order) {
	cookPeriod := getCookPeriod()
	log.Println("cooking orderId:", order.OrderId)
	time.Sleep(time.Duration(cookPeriod) * time.Second)
	log.Println("cooked orderId:", order.OrderId, "after", cookPeriod, "seconds")
	order.Status = "DONE"
}

func getCookPeriod() int {
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	return randomizer.Intn(viper.GetInt("app.cook_period"))
}
