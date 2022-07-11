package function

import (
	"BackEnd/config"
	"BackEnd/expen"
	"BackEnd/service"
	"BackEnd/service/ForumListTable"
	"BackEnd/service/UserAccountTable"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"time"
)

// ArticleContent 文章返回的接口数据
type ArticleContent struct {
	ID          string               `json:"ID" `         // 文章ID
	ArticleData ArticleRedis         `json:"ArticleData"` // redis储存的各样数据
	Article     ArticleElasticsearch `json:"Article"`     // 文章存在elasticsearch的内容
}

// ArticleRedis 文章储存在redis里的内容
type ArticleRedis struct {
	Title             string `json:"Title"`             // 文章标题
	PageViews         int    `json:"PageViews"`         // 浏览量
	Likes             int    `json:"Likes"`             // 点赞量
	AuthorUID         string `json:"AuthorUID"`         // 作者UID
	ViewPermissions   bool   `json:"ViewPermissions"`   // 是否需要登录查看
	ClassificationUID string `json:"ClassificationUID"` // 贴吧的UID
}

// ArticleElasticsearch 定义储存在elasticsearch 索引的文章储存内容
type ArticleElasticsearch struct {
	Description        string    `json:"Description"`        // 简略文章内容
	Img                string    `json:"Img"`                // 图片第一个地址
	Content            string    `json:"Content"`            // 文章内容
	AuthorName         string    `json:"AuthorName"`         // 创作者名称
	ClassificationName string    `json:"ClassificationName"` // 文章对应分类贴吧名称
	AuthorIMG          string    `json:"AuthorIMG"`          // 创作者头像地址
	ClassificationIMG  string    `json:"ClassificationIMG"`  // 文章对应分类贴吧头像地址
	CreateTime         time.Time `json:"CreateTime"`         // 文章创建时间
	Keywords           string    `json:"Keywords"`           // 关键词
}

// ArticleInterface 创建文章的对象接口 https://juejin.cn/post/7015426086699794469 为什么要使用接口形式
type ArticleInterface interface {
	Init() bool              // 插入总功能
	GetVal() error           // 赋值所有数据
	InsertRedis() error      // 将杂项数据插入redis
	AddPageViews(IP string)  // 增加浏览量
	AddLikes(uid string)     // 增加点赞接口
	InsertArticle() error    // 文章存储进入Elasticsearch
	ElasticsearchGet() error // 获取Elasticsearch里的值
	MysqlGetUser()           // 获取mysql中储存的用户头像和名称
	MysqlGetForum()          // 获取mysql储存的论坛名字和头像
}

// GetVal 赋值所有数据
func (Article *ArticleContent) GetVal() error {
	// 读取redis里面的内容
	ArticleAll := expen.HashReadAll(config.RedisArticle, Article.ID)
	// 将储存在hash表的数据赋值给结构体
	if err := mapstructure.WeakDecode(ArticleAll, &Article.ArticleData); err != nil {
		logrus.Error("转换Redis中的文章数据失败:", err)
		return err
	}
	// 将存进elasticsearch 的文章内容赋值
	if err := Article.ElasticsearchGet(); err != nil {
		logrus.Error("转换Elasticsearch中的文章数据失败:", err)
		return err
	}
	// 将两个字段的内容值转换
	return nil
}

// ElasticsearchGet  获取elasticsearch 的索引内容
func (Article *ArticleContent) ElasticsearchGet() error {
	return expen.ElasticsearchGet(Article.ArticleData.ClassificationUID, Article.ID, &Article.Article)
}

// InsertRedis 将杂项数据插入redis
func (Article *ArticleContent) InsertRedis() error {
	// 改为批量插入
	if _, err := config.RedisArticle.HMSet(Article.ID, map[string]interface{}{
		"Title":             Article.ArticleData.Title,
		"PageViews":         Article.ArticleData.PageViews,
		"Likes":             Article.ArticleData.Likes,
		"AuthorUID":         Article.ArticleData.AuthorUID,
		"ViewPermissions":   Article.ArticleData.ViewPermissions,
		"ClassificationUID": Article.ArticleData.ClassificationUID,
	}).Result(); err != nil {
		logrus.Info("批量文章插入信息失败：", err)
		return err
	}
	return nil
}

// InsertArticle  将文章内容插入elasticsearch中
func (Article *ArticleContent) InsertArticle() error {
	return expen.ElasticsearchInsert(Article.ArticleData.ClassificationUID, Article.ID, Article.Article)
}

// AddPageViews 增加浏览量
func (Article *ArticleContent) AddPageViews(IP string) {
	// 查询文章Ip里面是否有相同的IP，单个文章相同IP不作数
	if Null, _ := expen.HaseRead(config.RedisComment, Article.ID, IP); !Null {
		// 插入IP和浏览量
		expen.HaseSet(config.RedisComment, Article.ID, IP, IP)
		expen.HashInsertAdd(config.RedisArticle, Article.ID, "PageViews")
	}

}

// AddLikes 增加点赞接口
func (Article *ArticleContent) AddLikes(uid string) {
	// 查询文章Ip里面是否有相同用户的UID，单个文章相同UID已点赞就无法点赞了
	if Null, _ := expen.HaseRead(config.RedisComment, Article.ID, uid); !Null {
		// 插入uid和点赞数
		expen.HaseSet(config.RedisComment, Article.ID, uid, "true")
		expen.HashInsertAdd(config.RedisArticle, Article.ID, "Likes")
	}
}

// LikesBool 是否点赞接口
func (Article *ArticleContent) LikesBool(uid string) bool {
	Null, _ := expen.HaseRead(config.RedisComment, Article.ID, uid)
	return Null
}

// Init 插入文章到各组件的初始化
func (Article *ArticleContent) Init() bool {
	// 首先是插入redis
	Article.Article.CreateTime = time.Now()
	if err := Article.InsertRedis(); err != nil {
		return false
	}
	// 获取到UID之后直接Mysql获取内容
	Article.MysqlGetUser()
	Article.MysqlGetForum()
	Article.Article.Description = expen.StringsP(config.LuteEngine.MarkdownStr(Article.ArticleData.Title, Article.Article.Content))
	// 插入Elasticsearch
	if err := Article.InsertArticle(); err != nil {
		logrus.Info("插入文章数据失败:", err)
		return false
	}

	// 将文章部分数据插入文章表
	if err := Article.MysqlArticle(); err != nil {
		logrus.Error("文章插入mysql失败:", err)
		return false
	}

	return true
}

// MysqlGetUser 获取mysql中储存的用户头像和名称
func (Article *ArticleContent) MysqlGetUser() {
	UserAccount := UserAccountTable.UserIMGAndName(Article.ArticleData.AuthorUID)
	Article.Article.AuthorIMG = UserAccount.Img
	Article.Article.AuthorName = UserAccount.UserName
}

// MysqlGetForum 获取mysql储存的论坛名字和头像
func (Article *ArticleContent) MysqlGetForum() {
	ForumList := ForumListTable.ForumIMGAndName(Article.ArticleData.ClassificationUID)
	Article.Article.ClassificationIMG = ForumList.ImgURL
	Article.Article.ClassificationName = ForumList.Name
}

// MysqlArticle 数据库增加文章大致内容内容
func (Article *ArticleContent) MysqlArticle() error {
	var table service.Table
	table.Title = Article.ArticleData.Title
	table.AuthorName = Article.Article.AuthorName
	table.Uid = Article.ID
	table.Content = expen.StringsP(config.LuteEngine.MarkdownStr(Article.ArticleData.Title, Article.Article.Content))
	table.AuthImg = Article.Article.AuthorIMG
	table.Img = Article.Article.Img
	return service.CreateArticleContent(table, Article.ArticleData.ClassificationUID)
}

// 数据库增加文章评论表
