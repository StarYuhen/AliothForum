package use

import (
	"BackEnd/config"
	"BackEnd/expen"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

// JwtPO 自定义jwt参数
type jwtPO struct {
	User      string `json:"user"`
	Uid       string `json:"uid"`
	Authority string `json:"authority"`
	jwt.StandardClaims
}

// TokenExpireDuration 设置为100天，一次用户登录有效期
const TokenExpireDuration = time.Hour * 2400

var MySecret = []byte(config.StoreConfig.WebEncrypt.JWT)

// GenToken 生成JWT
func GenToken(user, uid, Authority string) (string, error) {
	// 创建一个我们自己的声明
	lie := jwtPO{
		user,
		uid,
		Authority,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "StarYuhen",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, lie)
	tokens, _ := token.SignedString(MySecret)
	logrus.Info("加密过后的jwt---->", tokens)
	return tokens, nil
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*jwtPO, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &jwtPO{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 将解析后的内容返回
	if clime, ok := token.Claims.(*jwtPO); ok && token.Valid {
		// logrus.Info("解析后的jwt内容---->", clime)
		return clime, nil
	}

	return nil, errors.New("invalid token")
}

// AuthMiddlewareJWT  基于JWT的认证中间件
func AuthMiddlewareJWT() func(c *gin.Context) {
	return func(ctx *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			// jwt不存在
			ctx.JSON(http.StatusOK, expen.ParameterErrorFun("未登录账号，请登录"))
			logrus.Info("请求头中auth为空")
			ctx.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			// jwt没有加上Bearer
			ctx.JSON(http.StatusOK, expen.ParameterErrorFun("登录账号错误，请重新登录"))
			ctx.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := ParseToken(parts[1])
		if err != nil {
			// jwt过期了或无效
			ctx.JSON(http.StatusOK, expen.ParameterErrorFun("账号已过期，请重新登录"))
			ctx.Abort()
			logrus.Info("无效的Token")
			return
		}

		// // 检查redis是否有相同redis--不允许同时登陆
		// if bo, jwts := redisDatabases.HaseRead("Jwt", mc.Uid); !bo || jwts != parts[1] {
		// 	ctx.JSON(http.StatusOK, InternalErrorFun("账号已在异地登录或者未登录过"))
		// 	logrus.Info("该jwt并未被使用:", parts[1])
		// 	ctx.Abort()
		// 	return
		// }

		// 将当前请求的user和uid信息保存到请求的上下文c上
		ctx.Set("user", mc.User)
		ctx.Set("uid", mc.Uid)
		ctx.Set("Authority", mc.Authority)
		ctx.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
