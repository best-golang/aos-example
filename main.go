package main

import (
	"github.com/aos-stack/aos"
	"github.com/gin-gonic/gin"
	"./app"
)
func main() {
	aos.RegisterRouter(func (engine *gin.Engine) {
		engine.GET("/", app.NewHelloWorldController().Index)
	})
	aos.Run()
}