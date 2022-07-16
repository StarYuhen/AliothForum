package test

import (
	"BackEnd/config"
	"BackEnd/expen"
	"BackEnd/server/api/function"
	"BackEnd/service"
	"BackEnd/service/UserAccountTable"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

// 插入文章
func TestArticleInsert(receiver *testing.T) {
	Article := new(function.ArticleContent)
	Article.ID = "Test-Article"
	// Article.Article = "/Test-Article.md"
	Article.InsertRedis()
}

// 获取存储的文章内容
func TestArticleRead(t *testing.T) {
	Article := new(function.ArticleContent)
	Article.ID = "Test-Article"
	logrus.Info(Article.GetVal())
	logrus.Info(Article)
	// 增加字段
	expen.HashInsertAdd(config.RedisArticle, "Test-Article", "PageViews")
}

// 插入es
func TestElasticsearchInsert(t *testing.T) {
	var ArticleContent function.ArticleContent
	ArticleContent.ID = "test sT_grg ?tsdf d"
	ArticleContent.ArticleData.Title = "4354-5676-tr-345"
	ArticleContent.Article.Content = "### 测试下elasticsearch 是否能够正常使用 "
	ArticleContent.Article.CreateTime = time.Now()
	logrus.Info(ArticleContent.InsertArticle())
}

// 获取es 内容
func TestElasticsearchGet(t *testing.T) {
	var ArticleContent function.ArticleContent
	ArticleContent.ID = "1"
	ArticleContent.ArticleData.Title = "测试一下"
	ArticleContent.ElasticsearchGet()
	logrus.Info(ArticleContent.Article)
}

func TestUSERName(t *testing.T) {
	logrus.Info(UserAccountTable.UserIMGAndName("88f50918-cde0-11ec-94fb-8c8caad3310e"))
}

// 创建文章储存表
func TestTable(t *testing.T) {
	expen.CreateArticle("88f50918-cde0-11ec-94fb-8c8caad3310e", service.Table{})
}

// 创建文章评论表
func TestTableComment(t *testing.T) {
	expen.CreateArticle("88f50918-cde0-11ec-94fb-8c8caad3310e_Comment", service.Comment{})
}

// 获取文章随机链接
func TestTableUREL(t *testing.T) {
	service.RandArticleURL("88f50918-cde0-11ec-94fb-8c8caad3310e")
	// 随机获取hash的内容
	logrus.Info(expen.HashRandomKey(config.RedisArticle))
}

// 获取储存的所有key
func TestRedisKeys(t *testing.T) {
	logrus.Info(expen.HashKeyS(config.RedisArticle))
}

// 插入评论
func TestComment(t *testing.T) {
	for i := 0; i < 5; i++ {
		logrus.Info(service.CreateComment(service.Comment{
			Uid:           "1657156821922318800",
			AuthorName:    "StarYuhen",
			AuthImg:       "https://www.yuhenm.com/usr/uploads/2022/07/2588829332.png",
			CommentOneUid: uuid.NewString(),
			CommentText:   "测试一级评论",
		}, "88f50918-cde0-11ec-94fb-8c8caad3310e_comment"))
	}
}
