package middlewares

import (
	"github.com/kataras/iris"
)

// ErrorHandler 通用错误处理
func ErrorHandler(ctx iris.Context) {
	defer func() {
		if err := recover(); err != nil {
			if ctx.GetStatusCode() == iris.StatusOK {
				ctx.StatusCode(iris.StatusInternalServerError)
			}
			ctx.JSON(err)
		}
	}()

	ctx.Next()
}
