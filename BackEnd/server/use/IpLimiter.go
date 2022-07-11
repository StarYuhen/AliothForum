package use

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"time"
)

// NewRateLimiterMiddleware 限制每个ip的时间请求接口限制，使用滑动窗口算法 https://github.com/imtoori/gin-redis-ip-limiter
func NewRateLimiterMiddleware(redisClient *redis.Client, key string, limit int, slidingWindow time.Duration) gin.HandlerFunc {
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprint("error init redis", err.Error()))
	}

	return func(c *gin.Context) {
		now := time.Now().UnixNano()
		userCntKey := fmt.Sprint(c.ClientIP(), ":", key)

		redisClient.ZRemRangeByScore(userCntKey,
			"0",
			fmt.Sprint(now-(slidingWindow.Nanoseconds()))).Result()

		reqs, _ := redisClient.ZRange(userCntKey, 0, -1).Result()

		if len(reqs) >= limit {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code": http.StatusTooManyRequests,
				"msg":  "请求超过限制次数太多啦，等一秒哦~",
			})
			return
		}

		c.Next()
		redisClient.ZAddNX(userCntKey, redis.Z{Score: float64(now), Member: float64(now)})
		redisClient.Expire(userCntKey, slidingWindow)
	}

}
