package src

import (
	"github.com/ghoshRitesh12/gojo/src/config"
	"github.com/ghoshRitesh12/gojo/src/providers/anicrush"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	app := gin.Default()

	app.Group("/anicrush").
		GET("/home", anicrush.HomeHandler)

	app.Use(config.NotFoundHandler)

	return app
}
