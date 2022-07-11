package config

import (
	"github.com/panjf2000/ants"
	"github.com/sirupsen/logrus"
)

// AntsPoolNew  全局对象协程池 -- 将所有协程任务写入协程池中 减少资源消耗
func AntsPoolNew() *ants.Pool {
	p, err := ants.NewPool(StoreConfig.WebConfig.Ant, ants.WithPreAlloc(true)) // 2000w协程池容量--预先分配内存减少使用
	if err != nil {
		logrus.Error("协程池初始化失败:", err)
		panic(err)
	}
	logrus.Info("已初始化协程池")
	return p
}
