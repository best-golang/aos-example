package main

import (
	container "github.com/aos-stack/aos-example/app/bind-service"
	router "github.com/aos-stack/aos-example/app/routers"
	"github.com/aos-stack/aos-example/env"
	"github.com/aos-stack/aos"
	"github.com/aos-stack/aos/interfaces"
	"github.com/gin-gonic/gin"
)

func main() {
	interfaces.AddCommand("EnvExampleCommand", env.EnvExampleCommand{})

	aos.RegisterRouter(func(engine *gin.Engine) {
		container.GetContainer()
		router.InitRouter(engine)
		//
	})
	aos.Run()
}
