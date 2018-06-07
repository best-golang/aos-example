package main

import (
	container "aos-example/app/bind-service"
	router "aos-example/app/routers"
	"aos-example/env"

	"github.com/aos-stack/aos"
	"github.com/aos-stack/aos/interfaces"
	"github.com/gin-gonic/gin"
)

func main() {
	interfaces.AddCommand("EnvExampleCommand", env.EnvExampleCommand{})

	aos.RegisterRouter(func(engine *gin.Engine) {
		container.GetContainer()
		router.InitRouter(engine)
	})
	aos.Run()
}
