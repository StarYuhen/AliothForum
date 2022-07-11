package use

import (
	"BackEnd/expen"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Limiter 使用ToolBooth 限制浏览中间件
func Limiter(limit *limiter.Limiter) gin.HandlerFunc {
	return func(context *gin.Context) {
		// 当超过初始化配置的限制时，就返回超过的限制的信息
		if HttpError := tollbooth.LimitByRequest(limit, context.Writer, context.Request); HttpError != nil {
			context.JSON(http.StatusTooManyRequests, expen.ParameterErrorFun("请求超过限制次数太多啦，等一秒哦~"))
			context.Abort()
		}
		context.Next()
	}
}
