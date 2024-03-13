package home

import (
	"log"

	"github.com/ghoshRitesh12/gojo/src/config"
	"github.com/gin-gonic/gin"
)

var Handler gin.HandlerFunc = func(ctx *gin.Context) {
	data, err := scrapeHomePage()
	if err != nil {
		log.Println(err)

		if config.IsHttpError(err) {
			config.HandleErr(ctx, err.(*config.HttpError))
			return
		}

		config.HandleErr(ctx, config.NewHttpError(500))
		return
	}

	ctx.JSON(200, data)
}
