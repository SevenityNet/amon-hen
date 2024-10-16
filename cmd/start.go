package cmd

import (
	"amonhen/controllers"
	"amonhen/middlewares"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func Execute() {
	app := gin.Default()
	app.Use(middlewares.CORS())
	app.Use(middlewares.Recovery())
	app.Use(middlewares.ErrorHandler())

	controllers.Register(app.RouterGroup)

	if err := app.Run(); err != nil {
		log.Fatal("😭 Failed to start Amon Hen.", err)
	}
}
