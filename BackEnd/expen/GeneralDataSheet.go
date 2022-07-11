package expen

import (
	"BackEnd/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// 用于创建通用的数据表，如一个子论坛所有文章需要在一个表 General Data Sheet

// CreateArticle 创建通用储存表
func CreateArticle(TableName string, Table interface{}) bool {
	// 先通过结构体创建一个临时表
	// 检测是否存在那张表
	if !config.MysqlURL.Migrator().HasTable(TableName) {
		if err := config.MysqlURL.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&Table); err != nil {
			logrus.Error("创建论坛文章表失败:", err)
			return false
		}
		if err := config.MysqlURL.Migrator().RenameTable(&Table, TableName); err != nil {
			logrus.Error("重命名论坛文章表失败:", err)
			return false
		}
		return true
	}

	return true
}

// Paginate 分页器计算函数
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch {
		case pageSize > 100:
			pageSize = 20
		case pageSize <= 0:
			pageSize = 20
		}
		return db.Offset(page).Limit(pageSize)
	}
}
