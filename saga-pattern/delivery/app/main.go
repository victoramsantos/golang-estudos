package main

import (
	"os"
	"sagapattern/delivery/repository"
	"sagapattern/delivery/usecase"

	"github.com/spf13/viper"
)

func main() {
	barCounterRepository := repository.NewBarCounterRepository()
	usecase := usecase.NewDeliveryUsecase(
		barCounterRepository,
	)
	usecase.Delivery()
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
