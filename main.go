package main

import (
	"aos-example/app"
	"aos-example/env"

	"github.com/aos-stack/aos"
	"github.com/aos-stack/aos/interfaces"
	"github.com/gin-gonic/gin"
)

func main() {

	interfaces.AddCommand("EnvExampleCommand", env.EnvExampleCommand{})

	aos.RegisterRouter(func(engine *gin.Engine) {
		engine.GET("/", app.NewHelloWorldController().Index)
	})
	aos.Run()
}
