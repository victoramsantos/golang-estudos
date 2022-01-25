package usecase

import (
	"log"
	"math/rand"
	"sagapattern/delivery/domain"
	"time"

	"github.com/spf13/viper"
)

type deliveryUsecase struct {
	barCounterRepository domain.BarCounterRepository
}

func NewDeliveryUsecase(barCounterRepository domain.BarCounterRepository) domain.DeliveryUsecase {
	return &deliveryUsecase{
		barCounterRepository: barCounterRepository,
	}
}

func (usecase *deliveryUsecase) Delivery() {
	for {
		order, err := usecase.barCounterRepository.Delivery()
		if err != nil {
			log.Println("No message found, sleeping", viper.GetInt("app.sleep_period"), "seconds")
			time.Sleep(time.Duration(viper.GetInt("app.sleep_period")) * time.Second)
			continue
		}
		//Should be a go func
		delivery(order)
	}
}

func delivery(order *domain.Order) {
	cookPeriod := getDeliveryPeriod()
	log.Println("delivery orderId:", order.OrderId)
	time.Sleep(time.Duration(cookPeriod) * time.Second)
	log.Println("delivery orderId:", order.OrderId, "after", cookPeriod, "seconds")
	order.Status = "DELIVERED"
}

func getDeliveryPeriod() int {
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	return randomizer.Intn(viper.GetInt("app.delivery_period"))
}
