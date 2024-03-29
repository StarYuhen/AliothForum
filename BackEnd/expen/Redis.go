package expen

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"time"
)

// StringSet 设置redis的字符串及过期时间
func StringSet(client *redis.Client, key, value string, times time.Duration) bool {
	// 统一过期时间60s
	err := client.Set(context.Background(), key, value, times).Err()
	if err != nil {
		logrus.Info("插入redis失败，请刷新")
		return false
	}
	return true
}

// StringRead 获取redis字符串的值
func StringRead(client *redis.Client, key string) (bool, string) {
	err := client.Get(context.Background(), key).Val()
	if err == "" {
		logrus.Info("查询redis结果为空，请刷新")
		return false, ""
	}
	return true, err
}

// HaseSet 向hash表插入数据 写入已有key会自动更新已有键的值
func HaseSet(client *redis.Client, key string, findKey string, value interface{}) error {
	// 向hash数据库读取数据
	_, err := client.HSet(context.Background(), key, findKey, value).Result()
	if err != nil {
		logrus.Error("插入redis哈希表失败--->", err)
		return err
	}
	return nil
}

// HaseRead 向指定hash表获取对应key的value
func HaseRead(client *redis.Client, key string, findKey string) (bool, string) {
	// 向hash数据库读取数据
	res, err := client.HMGet(context.Background(), key, findKey).Result()
	if err != nil {
		logrus.Error("读取redis哈希表失败--->", err)
		return false, ""
	}
	if res[0] == nil {
		return false, ""
	}
	return true, res[0].(string)
}

// HashReadAll 获取hash表所有元素
func HashReadAll(client *redis.Client, key string) map[string]string {
	return client.HGetAll(context.Background(), key).Val()
}

// HashReadAllKey 获取hash表的所有key
func HashReadAllKey(client *redis.Client, key string) []string {
	return client.HKeys(context.Background(), key).Val()
}

// HashInsertAdd 实现hash表的自增
func HashInsertAdd(client *redis.Client, key string, find string) int64 {
	return client.HIncrBy(context.Background(), key, find, 1).Val()
}

// HashRandomKey 随机获取一个Hash表Key
func HashRandomKey(client *redis.Client) string {
	return client.RandomKey(context.Background()).Val()
}

// HashKeyS 获取hash表所有的key
func HashKeyS(client *redis.Client) []string {
	return client.Keys(context.Background(), "*").Val()
}
