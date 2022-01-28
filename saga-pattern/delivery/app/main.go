package main

import (
	"log"
	"os"
	"sagapattern/delivery/repository"
	"sagapattern/delivery/usecase"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	e := echo.New()

	initPrometheus(e)
	go initDeliveryService()

	log.Fatal(e.Start(":" + viper.GetString("app.server.port")))

}

func initPrometheus(e *echo.Echo) {
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)
}

func initDeliveryService() {
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
