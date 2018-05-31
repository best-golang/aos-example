package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloWorldController struct {

}

func (this *HelloWorldController)Index(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}

func NewHelloWorldController() *HelloWorldController {
	return &HelloWorldController{}
}