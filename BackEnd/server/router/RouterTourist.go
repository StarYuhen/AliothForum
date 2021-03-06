package router

import (
	"BackEnd/server/api/tourist"
	"BackEnd/server/use"
	"github.com/gin-gonic/gin"
)

// TouristRouter 游客状态可以访问的接口
func TouristRouter(engine *gin.RouterGroup) {
	// 请求验证码接口
	engine.GET("/captcha", tourist.Captcha)
	// 注册账号接口
	engine.POST("/registerAccount", tourist.RegisterAccount)
	// 登录账号接口
	engine.POST("/loginAccount", tourist.LoginAccount)
	// 请求邮件验证码
	engine.GET("/registerAccountMail", use.MajorIPSleep(), tourist.RegisterAccountMail)
	// 请求文章内容--接口请求的结果
	engine.GET("/article/:ID", tourist.GetArticleContent)
}
