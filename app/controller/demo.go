package controller

import (
	"fmt"
	"net/http"

	"aos-example/demo-ddd/service"

	aosContainer "github.com/aos-stack/aos/container"
	"github.com/gin-gonic/gin"
)

// DemoApi ...
type DemoApi struct {
	DemoService service.DemoService
}

// NewDemoApi ...
func NewDemoApi(demoService service.DemoService) *DemoApi {
	return &DemoApi{
		DemoService: demoService,
	}
}

// Index ...
func (myc *DemoApi) Index(c *gin.Context) {
	logger := aosContainer.GetLogger("logger")

	logger.Info("hhhhhhhhhhh")

	fmt.Println(aosContainer.ContainerGet("consul"))
	c.String(http.StatusOK, "hello world")
}

// Index2 ...
func (myc *DemoApi) Index2(c *gin.Context) {
	myc.DemoService.List(3)
	c.String(http.StatusOK, "hello world2")
}
