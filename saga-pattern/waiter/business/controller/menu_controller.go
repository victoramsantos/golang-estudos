package controller

import (
	"net/http"
	"sagapattern/waiter/domain"

	"github.com/labstack/echo"
)

type menuController struct {
	usecase domain.MenuUsecase
}

func InitMenuController(e *echo.Echo, usecase domain.MenuUsecase) {
	handler := &menuController{
		usecase: usecase,
	}

	e.GET("/menu", handler.ShowMenu)
}

func (ctl *menuController) ShowMenu(ctx echo.Context) error {
	menu := ctl.usecase.ShowMenu()
	return ctx.JSON(http.StatusOK, menu)
}
