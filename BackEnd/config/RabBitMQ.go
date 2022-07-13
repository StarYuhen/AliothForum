package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

// RabBitMQ 初始化的消息队列引擎 教程 https://www.liwenzhou.com/posts/Go/go_rabbitmq_tutorials_01/
func RabBitMQ() *amqp.Connection {
	URL := fmt.Sprintf("amqp://%s:%s@%s/%s",
		StoreConfig.RabBitMQ.User,
		StoreConfig.RabBitMQ.PassWord,
		StoreConfig.RabBitMQ.Addr,
		StoreConfig.RabBitMQ.VirtualHost,
	)
	client, err := amqp.Dial(URL)
	if err != nil {
		logrus.Error("启动RabBitMQ消息队列组件失败：", err)
		panic(err)
	}
	logrus.Info("初始化RabBitMQ消息队列组件成功")
	return client
}
