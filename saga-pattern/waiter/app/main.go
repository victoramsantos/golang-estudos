package main

import (
	"log"
	"sagapattern/waiter/business/controller"
	"sagapattern/waiter/business/repository"
	"sagapattern/waiter/business/usecase"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	repository := repository.NewOrderRepository()
	usecase := usecase.NewOrderUsecase(repository)
	controller.NewOrderController(e, usecase)

	log.Fatal(e.Start(":8080"))

}
