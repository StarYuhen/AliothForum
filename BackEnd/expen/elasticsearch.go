package expen

import (
	"BackEnd/config"
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
)

// ElasticsearchInsert 用于储存内容
func ElasticsearchInsert(index, ID string, Json interface{}) error {
	ctx := context.Background()
	if _, err := config.ElasticsearchEngine.Index().
		Index(index).   // 设置索引名称
		Id(ID).         // 设置文档id
		BodyJson(Json). // 指定前面声明struct对象
		Do(ctx); err != nil {
		logrus.Error("elasticsearch储存文章失败:", err)
		return err
	}
	return nil
}

// ElasticsearchGet 用于读取储存的内容
func ElasticsearchGet(index, ID string, Json interface{}) error {
	ctx := context.Background()
	get, err := config.ElasticsearchEngine.Get().
		Index(index).
		Id(ID).
		Do(ctx)

	data, err := get.Source.MarshalJSON()

	if err := json.Unmarshal(data, Json); err != nil {
		logrus.Info("json 转换失败:", err)
		return err
	}

	if err != nil {
		logrus.Error("elasticsearch转换数据失败:", err)
		return err
	}

	return nil
}
