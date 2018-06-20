package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/aos-stack/aos-example/demo-ddd/service"
	aosContainer "github.com/aos-stack/aos/container"
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

	session := sessions.Default(c)
	var count int
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count++
	}
	session.Set("count", count)
	session.Save()
	fmt.Println(count)
	c.String(http.StatusOK, "hello world no ddd")
}

// Index2 ...
func (myc *DemoApi) Index2(c *gin.Context) {
	c.String(http.StatusOK, "hello world2 no ddd")
}
