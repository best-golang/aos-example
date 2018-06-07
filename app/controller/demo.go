package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DemoApi ...
type DemoApi struct {
}

// NewDemoApi ...
func NewDemoApi() *DemoApi {
	return &DemoApi{}
}

// Index ...
func (myc *DemoApi) Index(c *gin.Context) {
	c.String(http.StatusOK, "hello world no ddd")
}

// Index2 ...
func (myc *DemoApi) Index2(c *gin.Context) {
	c.String(http.StatusOK, "hello world2 no ddd")
}
