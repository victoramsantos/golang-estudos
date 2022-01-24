package main

import (
	"log"
	"os"
	"sagapattern/waiter/controller"
	"sagapattern/waiter/repository"
	"sagapattern/waiter/usecase"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func main() {
	e := echo.New()

	initOrderService(e)
	initMenuService(e)

	log.Fatal(e.Start(":" + viper.GetString("app.server.port")))
}

func initOrderService(e *echo.Echo) {
	repository := repository.NewOrderRepository()
	usecase := usecase.NewOrderUsecase(repository)
	controller.InitOrderController(e, usecase)
}

func initMenuService(e *echo.Echo) {
	repository := repository.NewMenuRepository()
	usecase := usecase.NewMenuUsecase(repository)
	controller.InitMenuController(e, usecase)
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
