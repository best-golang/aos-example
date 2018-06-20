package router

import (
	"github.com/aos-stack/aos-example/app/controller"

	aosContainer "github.com/aos-stack/aos/container"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {

	demoDDDApi := aosContainer.ContainerGet("demoDDDApi").(*controller.DemoDDDApi)

	demoApi := controller.NewDemoApi()
	engine.GET("/", demoApi.Index)
	engine.GET("/test", demoApi.Index2)

	engine.GET("/hasddd", demoDDDApi.Index)
	engine.GET("/hasddd/test", demoDDDApi.Index2)
}
