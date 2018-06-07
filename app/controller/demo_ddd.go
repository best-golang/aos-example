package controller

import (
	"fmt"
	"net/http"

	"aos-example/demo-ddd/service"

	aosContainer "github.com/aos-stack/aos/container"
	"github.com/gin-gonic/gin"
)

// DemoDDDApi ...
type DemoDDDApi struct {
	DemoService service.DemoService
}

// NewDemoApi ...
func NewDemoDDDApi(demoService service.DemoService) *DemoDDDApi {
	return &DemoDDDApi{
		DemoService: demoService,
	}
}

// Index ...
func (myc *DemoDDDApi) Index(c *gin.Context) {
	logger := aosContainer.GetLogger("logger")

	logger.Info("hhhhhhhhhhh")

	fmt.Println(aosContainer.ContainerGet("consul"))
	c.String(http.StatusOK, "hello world has ddd")
}

// Index2 ...
func (myc *DemoDDDApi) Index2(c *gin.Context) {
	myc.DemoService.List(3)
	c.String(http.StatusOK, "hello world2 has ddd")
}
