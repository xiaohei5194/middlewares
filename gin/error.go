package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler gin
func ErrorHandler(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if ctx.Writer.Status() == http.StatusOK {
				ctx.Status(http.StatusInternalServerError)
			}
			ctx.JSON(ctx.Writer.Status(), err)
		}
	}()

	ctx.Next()
}
