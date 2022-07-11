package use

import (
	"BackEnd/config"
	"BackEnd/expen"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// DecryptAPI ReptileBool 加密接口请求
/*
	// 如何实现防止爬虫？
	// ip接口限制已有中间件
	// 于是我想着在客户端每次请求都会随机一串加密数据，而这一串加密数据每次都能正常使用
	//
	// 采用服务器返回一个公钥，使用请求接口获取公钥  舍弃

	前端固定一套aes加密的公钥，然后校验前端的请求时间戳和设备浏览器生成指纹和MD5加密的StarYuhen，使用|进行分割
	而长时间每个请求都使用加密，导致业务出现资源消耗，所以每5分钟放开这个Sign，客户端5分钟之后才更换Sign
*/
func DecryptAPI() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 传递的加密数据
		Sign := context.Request.Header.Get("Sign")

		// 判断该IP是否进入黑名单
		if null, _ := expen.StringRead(config.RedisWebExpen, context.ClientIP()); null {
			context.JSON(http.StatusOK, expen.ParameterErrorFun("判断您当前使用IP有风险，1小时内无法访问网站，若有疑问可联系客服"))
			logrus.Info("请求接口IP为黑名单：", context.ClientIP())
			context.Abort()
			return
		}
		// 检查redis是否有这个Sign
		if Null, _ := expen.StringRead(config.RedisWebExpen, Sign); Null {
			// 进入则证明有
			context.Next()
			return
		}

		// 没有这个Sign就判断是否是正确的加密数据--一旦检查加密错误直接封锁Ip--加密时间戳无法对应上也直接封锁IP
		if str, err := expen.RSADecrypt([]byte(Sign), config.StoreConfig.WebEncrypt.RSAPrivate); err == nil {
			// 进入这里说明是解析未报错的

			SignString := strings.Split(str, "|")
			// 判断时间戳是否正确
			if strconv.FormatInt(time.Now().Unix(), 10) == SignString[0] {
				expen.StringSet(config.RedisWebExpen, Sign, "null", 5*time.Minute)
				logrus.Info("sign值验证成功")
				context.Next()
				return
			}

			logrus.Error("接口验证失败Sign：", str)
			context.JSON(http.StatusOK, expen.ParameterErrorFun("接口验证失败，请刷新网站"))
			context.Abort()
			return
		}

		// 解析报错直接拉入黑名单
		// TODO  关于IP问题，业务上线后 https://imlht.com/archives/248/
		expen.StringSet(config.RedisWebExpen, context.ClientIP(), "null", time.Hour)
		context.JSON(http.StatusOK, expen.ParameterErrorFun("判断您当前使用IP有风险，1小时内无法访问网站，若有疑问可联系客服"))
		context.Abort()
		return
	}
}
