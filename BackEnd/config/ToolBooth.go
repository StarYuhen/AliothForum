package config

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"time"
)

// TooBooth 用于函数初始化限流的配置内容
func TooBooth() *limiter.Limiter {
	lmt := tollbooth.NewLimiter(StoreConfig.WebConfig.TollBooth.TokenBuckets,
		&limiter.ExpirableOptions{DefaultExpirationTTL: StoreConfig.WebConfig.TollBooth.DefaultExpirationTTL * time.Hour})
	// 设置请求限制方式
	lmt.SetMethods(StoreConfig.WebConfig.TollBooth.Methods)
	lmt.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"})
	return lmt
}
