package controller

import (
	"net/http"
	"sagapattern/waiter/domain"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	validator "gopkg.in/go-playground/validator.v9"
)

type OrderController struct {
	OrderUsecase domain.OrderUsecase
}

func InitOrderController(e *echo.Echo, usecase domain.OrderUsecase) {
	handler := &OrderController{
		OrderUsecase: usecase,
	}

	e.POST("/order", handler.OrderRequest)
}

func (ctl *OrderController) OrderRequest(ctx echo.Context) error {
	var order domain.Order
	err := ctx.Bind(&order)

	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusUnprocessableEntity, ResponseError{err.Error()})
	}

	var ok bool
	if ok, err = isRequestValid(&order); !ok {
		log.Error(err.Error())
		return ctx.JSON(http.StatusBadRequest, ResponseError{err.Error()})
	}

	err = ctl.OrderUsecase.OrderRequest(&order)

	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, ResponseError{err.Error()})
	}
	return ctx.JSON(http.StatusCreated, order)
}

func isRequestValid(m *domain.Order) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
