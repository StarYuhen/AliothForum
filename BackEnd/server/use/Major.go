package use

import (
	"BackEnd/config"
	"BackEnd/expen"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// MajorIPSleep 有些重要接口需要设置时间限制,用ip做限制 一级重要资产接口
func MajorIPSleep() gin.HandlerFunc {
	return func(context *gin.Context) {
		ip := context.ClientIP()
		url := context.Request.URL.Path
		if Bool, src := expen.StringRead(config.RedisWebExpen, ip+"|"+url); !Bool || src != "true" {
			expen.StringSet(config.RedisWebExpen, ip+"|"+url, "true", time.Second*60)
		} else {
			context.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code": http.StatusTooManyRequests,
				"msg":  "请求超过限制，请等待片刻再来",
			})
			return
		}
	}
}

// MajorIPSleepTwo 二级重要资产接口
func MajorIPSleepTwo() gin.HandlerFunc {
	return func(context *gin.Context) {
		ip := context.ClientIP()
		url := context.Request.URL.Path
		if Bool, src := expen.StringRead(config.RedisWebExpen, ip+"|two|"+url); !Bool || src != "true" {
			expen.StringSet(config.RedisWebExpen, ip+"|two|"+url, "true", time.Second*30)
		} else {
			context.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code": http.StatusTooManyRequests,
				"msg":  "请求超过限制，请等待片刻再来",
			})
			return
		}
	}
}
