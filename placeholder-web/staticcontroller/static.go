package staticcontroller

import (
	"net/http"
	"placeholder/commons"
	"placeholder/service"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	r.GET("/static", get)
}

// GET /poem/static
func get(c *gin.Context) {
	var content = service.GetStaticContent()

	response := commons.Response{
		Status: 200,
		Body:   content.Content,
	}

	c.JSON(http.StatusOK, response)
}
