package config

import (
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

func ElasticsearchClient() *elastic.Client {
	client, err := elastic.NewClient(
		// 设置ES服务地址，支持多个地址
		elastic.SetURL("http://localhost:9200"),
		// 设置基于http base auth验证的账号和密码
		// elastic.SetBasicAuth("user", "secret"),
		// 启用gzip压缩
		elastic.SetGzip(true),
		// 发生连接失败，就用这个,将其设置为false
		elastic.SetSniff(false),
		// 设置监控检查时间间隔
		elastic.SetHealthcheckInterval(10*time.Second),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stderr, "", log.LstdFlags)),
	)

	if err != nil {
		// Handle error
		logrus.Error("初始化Elasticsearch失败: ", err)
		panic(err)
	} else {
		logrus.Info("初始化Elasticsearch成功")
	}
	return client
}
