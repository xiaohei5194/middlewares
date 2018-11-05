package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
@Time : 2018/11/5 18:34 
@Author : zhoushuai
@File : error
@Software: GoLand
*/

// ErrorHandler gin
func ErrorHandler(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
	}()
	ctx.Next()
}
