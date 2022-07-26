package expen

import (
	"BackEnd/config"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
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

// ElasticsearchGet 使用ID和分类UID读取储存的内容
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

// ElasticsearchGetArticleRandom  用于文章随机推荐
func ElasticsearchGetArticleRandom() (*elastic.SearchResult, error) {
	ctx := context.Background()
	// 初始化查询--随机获取结果 随机10个
	// script := elastic.NewScript("Math.random()")
	get, err := config.ElasticsearchEngine.Search("article").
		Size(2).
		// 返回json格式内容
		Pretty(true).
		Do(ctx)

	if err != nil {
		logrus.Error("查询随机文章错误:", err)
		return nil, err
	}
	return get, err
}
