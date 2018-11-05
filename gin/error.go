package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ErrorHandler gin
func ErrorHandler(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
	}()
	ctx.Next()
}
