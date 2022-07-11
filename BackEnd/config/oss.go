package config

import (
	"github.com/sirupsen/logrus"
	tx "github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
)

// 上传oss之前必须先行创建需要的目录
const path = "./"
const perm = 777

// oss初始化配置

func TxOSS() *tx.Client {
	u, _ := url.Parse(StoreConfig.OSS.TxConfig.OSSUrl)
	b := &tx.BaseURL{BucketURL: u}
	client := tx.NewClient(b, &http.Client{
		Transport: &tx.AuthorizationTransport{
			SecretID:  StoreConfig.OSS.TxConfig.SecretID,
			SecretKey: StoreConfig.OSS.TxConfig.SecretKey,
		},
	})
	logrus.Info("已初始化腾讯云对象储存")
	return client
}

// MkDir 创建文件夹
func MkDir(path string, perm os.FileMode) bool {
	err := os.MkdirAll(path, perm)
	return err == nil
}
