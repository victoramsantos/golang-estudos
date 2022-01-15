package dynamiccontroller

import (
	"net/http"
	"placeholder/commons"
	"placeholder/service"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	r.GET("/dynamic", get)
}

// GET /poem/dynamic
func get(c *gin.Context) {
	var content = service.GetDynamicContent()

	response := commons.Response{
		Status: 200,
		Body:   content.Content,
	}

	c.JSON(http.StatusOK, response)
}
