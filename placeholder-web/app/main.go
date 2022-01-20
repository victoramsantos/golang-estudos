package main

import (
	"log"
	poemController "placeholder/poem/controller"
	poemRepository "placeholder/poem/repository"
	poemUsecase "placeholder/poem/usecase"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	poemRepository := poemRepository.NewPoemRepository()
	poemUsecase := poemUsecase.NewPoemUsecase(poemRepository)
	poemController.NewPoemController(e, poemUsecase)

	log.Fatal(e.Start(":8080"))

}
