package config

import (
	"github.com/gin-gonic/gin"
)

var NotFoundHandler gin.HandlerFunc = func(ctx *gin.Context) {
	HandleErr(ctx, NewHttpError(404))
}
