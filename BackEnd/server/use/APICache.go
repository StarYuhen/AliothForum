package use

import (
	"BackEnd/config"
	"github.com/chenyahui/gin-cache"
	"github.com/gin-gonic/gin"
	"time"
)

// Cache 设置接口缓存中间件，方便服务器的请求并发,设置两小时，只有文章需要这么配置
func Cache() gin.HandlerFunc {
	return cache.CacheByRequestURI(config.CacheAPI, 2*time.Hour)
}
