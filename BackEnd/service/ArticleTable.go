package service

import (
	"BackEnd/config"
	"BackEnd/expen"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Table 通用文章储存表结构
type Table struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"` // 主键且自增
	Uid        string `json:"Uid,omitempty"`                                // 创建文章时生成的时间戳ID
	Title      string `json:"Title,omitempty" gorm:"unique"`                // 文章标题-并标记为唯一
	Content    string `json:"Content,omitempty"`                            // 文章内容 最多50个字
	AuthorName string `json:"AuthorName,omitempty"`                         // 作者昵称
	AuthImg    string `json:"AuthImg,omitempty"`                            // 作者头像
	Img        string `json:"Img,omitempty"`                                // 图片，当文章有图片时默认展示第一张图片
	gorm.Model `json:"Gorm.Model"`
}

// Comment 文章评论表--设计繁琐
/*
	1.每一个文章拥有着一个UID，他是所有评论对应的父UID
	2.一级评论指的是在用户评论的第一个
	2.二级评论指的是在一级评论里面评论了的内容
*/
type Comment struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"` // 主键且自增
	Uid           string // 文章对应UID
	AuthorName    string // 评论昵称
	AuthImg       string // 评论头像
	CommentOneUid string // 一级评论UID
	CommentTwoUid string // 二级评论UID
	CommentText   string // 评论内容
	gorm.Model
}

// RandArticleURL 随机的一个文章地址
func RandArticleURL(tableName string) {
	var table Table
	config.MysqlURL.Table(tableName).Select("uid").Order("rand()").Take(&table)
	logrus.Info(table)
}

// CreateArticleContent 插入文章数据
func CreateArticleContent(table Table, tableName string) error {
	return config.MysqlURL.Table(tableName).Create(&table).Error
}

// RandomArticle 获取随机推荐6个文章
func RandomArticle(tableName string) ([]Table, error) {
	var t []Table
	err := config.MysqlURL.Table(tableName).Select("distinct title,content").Where("id>=(?)",
		config.MysqlURL.Table(tableName).Select("FLOOR( MAX(id) * RAND())")).Limit(6).
		Find(&t).Error
	return t, err
}

// CreateComment 插入评论数据
func CreateComment(comment Comment, TableName string) error {
	return config.MysqlURL.Table(TableName).Create(&comment).Error
}

// PaginationCommentOne 一级评论分页器
func PaginationCommentOne(i int, m int, uid string, TableName string) ([]Comment, int64) {
	var comment []Comment
	if m-i >= 25 {
		return comment, 0
	}
	result := config.MysqlURL.Table(TableName).Where("uid=? and comment_two_uid is not null", uid).Scopes(expen.Paginate(i, m)).Find(&comment)
	return comment, result.RowsAffected
}
