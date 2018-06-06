package app

import (
	"net/http"

	aosContainer "github.com/aos-stack/aos/container"
	"github.com/gin-gonic/gin"
)

type HelloWorldController struct {
}

// Index ...
func (this *HelloWorldController) Index(c *gin.Context) {
	logger := aosContainer.GetLogger("logger")

	logger.Info("hhhhhhhhhhh")
	c.String(http.StatusOK, "hello world")
}

func NewHelloWorldController() *HelloWorldController {
	return &HelloWorldController{}
}
