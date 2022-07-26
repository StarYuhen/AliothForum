package config

import (
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

// 初始化redis链接 --用于储存web的杂项数据内容
func redisWebExpen() *redis.Client {
	logrus.Info("已初始化redis链接")
	return redis.NewClient(&redis.Options{
		Network:      "tcp",
		Addr:         StoreConfig.RedisConfig.Addr,
		Password:     StoreConfig.RedisConfig.Password,
		DB:           0,                                    // redis数据库index
		PoolSize:     StoreConfig.RedisConfig.PoolSize,     // redis链接池，默认是4倍cpu数，这里固定 用于协程链接
		MinIdleConns: StoreConfig.RedisConfig.MinIdleConns, // 初始规定的redis，维护，让其不少于这个数
	})
}

// 初始化redis链接 --用于储存数据库文章各项内容
func redisArticle() *redis.Client {
	return redis.NewClient(&redis.Options{
		Network:      "tcp",
		Addr:         StoreConfig.RedisConfig.Addr,
		Password:     StoreConfig.RedisConfig.Password,
		DB:           1,                                    // redis数据库index
		PoolSize:     StoreConfig.RedisConfig.PoolSize,     // redis链接池，默认是4倍cpu数，这里固定 用于协程链接
		MinIdleConns: StoreConfig.RedisConfig.MinIdleConns, // 初始规定的redis，维护，让其不少于这个数
	})
}

// 初始化redis链接 --用于储存数据库文章其余数据
func redisComment() *redis.Client {
	return redis.NewClient(&redis.Options{
		Network:      "tcp",
		Addr:         StoreConfig.RedisConfig.Addr,
		Password:     StoreConfig.RedisConfig.Password,
		DB:           2,                                    // redis数据库index
		PoolSize:     StoreConfig.RedisConfig.PoolSize,     // redis链接池，默认是4倍cpu数，这里固定 用于协程链接
		MinIdleConns: StoreConfig.RedisConfig.MinIdleConns, // 初始规定的redis，维护，让其不少于这个数
	})
}
