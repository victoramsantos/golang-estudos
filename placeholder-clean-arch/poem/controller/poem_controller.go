package controller

import (
	"net/http"
	"placeholder/domain"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	validator "gopkg.in/go-playground/validator.v9"
)

type poemController struct {
	poemUsecase domain.PoemUsecase
}

func NewPoemController(e *echo.Echo, poemUsecase domain.PoemUsecase) {
	handler := &poemController{
		poemUsecase: poemUsecase,
	}

	e.GET("/poem", handler.Fetch)
	e.POST("/poem", handler.Store)
	e.GET("/poem/random", handler.FetchRandom)
	e.GET("/poem/static", handler.FetchStatic)
}

func (ctl *poemController) Fetch(context echo.Context) error {
	poems := ctl.poemUsecase.FetchUsecase()
	return context.JSON(http.StatusOK, poems)
}

func (ctl *poemController) FetchRandom(context echo.Context) error {
	poem := ctl.poemUsecase.FetchRandomUsecase()
	return context.JSON(http.StatusOK, poem)
}

func (ctl *poemController) FetchStatic(context echo.Context) error {
	poem := ctl.poemUsecase.FetchStaticUsecase()
	return context.JSON(http.StatusOK, poem)
}

func (ctl *poemController) Store(context echo.Context) error {
	var poem domain.Poem
	err := context.Bind(&poem)
	if err != nil {
		log.Error(err.Error())
		return context.JSON(http.StatusUnprocessableEntity, ResponseError{err.Error()})
	}

	var ok bool
	if ok, err = isRequestValid(&poem); !ok {
		log.Error(err.Error())
		return context.JSON(http.StatusBadRequest, ResponseError{err.Error()})
	}

	ctl.poemUsecase.Store(&poem)
	return context.JSON(http.StatusCreated, poem)
}

func isRequestValid(m *domain.Poem) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
