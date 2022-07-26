package use

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// ReptileBool 判断是否是除搜索引擎的爬虫中间件
/*
	// 如何实现防止爬虫？
	// ip接口限制已有中间件
	// 于是我想着在客户端每次请求都会随机一串加密数据，而这一串加密数据每次都能正常使用
	//
	// 采用服务器返回一个公钥，使用请求接口获取公钥  舍弃

	前端固定一套aes加密的公钥，然后校验前端的请求时间戳和设备浏览器生成指纹和IP，使用|进行分割
	而长时间每个请求都使用加密，导致业务出现资源消耗，所以每5分钟放开这个Sign，客户端5分钟之后才更换Sign
*/
func ReptileBool() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 传递的加密数据
		// Sign := context.Request.Header.Get("Sign")
		// 存入Redis
	}
}

// DetectionCrawler 封装官方爬虫中间件
func DetectionCrawler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("Reptile", strings.Contains(context.GetHeader("User-Agent"), "compatible;"))
		// logrus.Info("当前设备UA:", context.GetHeader("User-Agent"))
		context.Next()
	}
}
