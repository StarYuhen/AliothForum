package server

import (
	"BackEnd/config"
	"BackEnd/server/api/tourist"
	"BackEnd/server/router"
	"BackEnd/server/use"
	"embed"
	cache "github.com/chenyahui/gin-cache"
	"github.com/danielkov/gin-helmet"
	"github.com/dvwright/xss-mw"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"html/template"
	"time"
)

//go:embed tmpl/*
var tmpl embed.FS

// GinServer 启动gin服务框架
func GinServer() {
	engine := gin.Default()
	// 绑定模板文件
	engine.SetHTMLTemplate(template.Must(template.New("").ParseFS(tmpl, "tmpl/*.tmpl")))
	engineCookie := cookie.NewStore([]byte("Yuhenm.com_StarYuhen"))
	var xssM xss.XssMw
	// 注册中间件
	// 绑定跨域中间件
	engine.Use(use.Cors())
	// 绑定cookie中间件
	engine.Use(sessions.Sessions("AliothForum", engineCookie))
	// 绑定消除xss攻击中间件
	engine.Use(xssM.RemoveXss())
	// 绑定简易安全中间件
	engine.Use(helmet.Default())
	// 增加接口并发限制中间件--不需要，自己重新写一个ip限制接口中间件
	// engine.Use(use.MaxAllowed(config.StoreConfig.WebConfig.MaxAllowed))
	// 启用gzip
	engine.Use(gzip.Gzip(gzip.DefaultCompression))
	// 增加ip规定时间限制 压测命令 yace -c 10000 -n 100 -u http://localhost:47/api/user/JwtDevice
	// 使用redis作为ip限制的速率限制器
	// engine.Use(use.NewRateLimiterMiddleware(config.RedisWebExpen,
	// 	config.StoreConfig.WebConfig.NewRateLimiterMiddleware.Key,
	// 	config.StoreConfig.WebConfig.NewRateLimiterMiddleware.Limit,
	// 	time.Duration(config.StoreConfig.WebConfig.NewRateLimiterMiddleware.Time)*time.Second))

	// 速率限制器中间件，令牌桶算法不适用redis作为限制
	engine.Use(use.Limiter(config.Limit))

	// // 启用加密请求
	// engine.Use(use.DecryptAPI())
	// // 部分重要接口需要这个中间件，请自添加
	// engine.Use(use.MajorIPSleep())
	// 检测是否是爬虫中间件
	engine.Use(use.DetectionCrawler())
	// 设置gin上传文件大小
	engine.MaxMultipartMemory = 3 << 20 // 3MB
	// 初始化路径
	Major := engine.Group("/api/major")     // 重要接口
	Tourist := engine.Group("/api/tourist") // 游客状态可以使用的接口
	User := engine.Group("/api/user")       // 不那么重要，但要使用jwt的接口

	// 分组绑定中间件
	// 绑定jwt请求校验
	User.Use(use.AuthMiddlewareJWT())
	Major.Use(use.AuthMiddlewareJWT())
	// 绑定路由驱动
	router.MajorRouter(Major)
	router.TouristRouter(Tourist)
	router.UserRouter(User)

	// 绑定静态资源文件夹 --必须是登录用户才能访问
	// 生成二维码储存地址
	User.Static("/QrCodeURL", config.StoreConfig.WebFile.QrCodeURL)
	// 文章Markdown文件储存地址
	User.Static("/MarkdownURL", config.StoreConfig.WebFile.MarkdownURL)
	// 请求文章内容--爬虫地图给予内容 设置接口缓存时间为一周
	engine.GET("/archives/:ID", cache.CacheByRequestURI(config.CacheAPI, 84*time.Hour), tourist.GetArticleContentReptile)

	// 正式生产项目使用RunTLS实现绑定接口https

	// 判断是否启动gin失败
	if err := engine.Run(config.StoreConfig.WebConfig.GinPort); err != nil {
		logrus.Info("启用Web服务失败，请检查错误！")
		panic(err)
	}

}
