package main

import (
	"log"
	"sagapattern/waiter/controller"
	"sagapattern/waiter/repository"
	"sagapattern/waiter/usecase"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	initOrderService(e)
	initMenuService(e)

	log.Fatal(e.Start(":8080"))
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
