package use

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cors 设置允许跨域请求 中间件
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		orgin := context.Request.Header.Get("Origin")
		if orgin != "" {
			// 接收客户端发送的origin （重要！）
			context.Writer.Header().Set("Access-Control-Allow-Origin", orgin)
			// 简易防止xss攻击
			context.Writer.Header().Set("Content-Type", "text/javascript")
			// 服务器支持的所有跨域请求的方法
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			// 允许跨域设置可以返回其他子段，可以自定义字段
			context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token, session, Content-Type,captcha,Sign")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			// 设置缓存时间
			context.Header("Access-Control-Max-Age", "172800")
			// 允许客户端传递校验信息比如 cookie (重要)
			context.Header("Access-Control-Allow-Credentials", "true")
		}

		// 允许类型校验
		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != any(nil) {
				fmt.Printf("Panic info is: %v\n", err)
			}
		}()

		context.Next()
	}
}
