package main

import (
	"os"
	"sagapattern/cook/repository"
	"sagapattern/cook/usecase"

	"github.com/spf13/viper"
)

func main() {
	orderRepository := repository.NewOrderRepository()
	barCounterRepository := repository.NewBarCounterRepository()
	usecase := usecase.NewCookUsecase(
		barCounterRepository,
		orderRepository,
	)
	usecase.Cook()
}

func init() {
	environment, isSet := os.LookupEnv("ENVIRONMENT")
	if !isSet {
		environment = "local"
	}
	viper.SetConfigName(environment)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
