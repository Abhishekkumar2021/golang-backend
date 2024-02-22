package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Logger(ctx *gin.Context) {
	// Do something here
	ctx.Next()
}
