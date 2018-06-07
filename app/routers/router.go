package router

import (
	"aos-example/app/controller"

	aosContainer "github.com/aos-stack/aos/container"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {

	// engine.GET("/", app.NewHelloWorldController().Index)
	// engine.GET("/test", app.NewHelloWorldController().Index2)
	demoApi := aosContainer.ContainerGet("demoApi").(*controller.DemoApi)

	engine.GET("/", demoApi.Index)
	engine.GET("/test", demoApi.Index2)

}
