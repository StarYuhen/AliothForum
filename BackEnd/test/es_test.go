package test

import (
	"BackEnd/expen"
	"BackEnd/server/api/function"
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

func TestEsRandom(t *testing.T) {
	var f function.ArticleElasticsearch
	get, _ := expen.ElasticsearchGetArticleRandom()
	for _, v := range get.Each(reflect.TypeOf(f)) {
		logrus.Info(v.(function.ArticleElasticsearch))
	}
}
