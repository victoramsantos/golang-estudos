package main

import (
	"placeholder/dynamiccontroller"
	"placeholder/staticcontroller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	dynamiccontroller.Setup(router)
	staticcontroller.Setup(router)

	router.Run()
}
