package router

import (
	"BackEnd/server/api/major"
	"github.com/gin-gonic/gin"
)

func MajorRouter(engine *gin.RouterGroup) {
	// 新建文章
	engine.POST("/insertArticle", major.InsertArticle)
	// 创建评论评论
	engine.POST("/InsertComment", major.InsertComment)
	// 创建论坛
	engine.POST("/CreateClassification", major.CreateClassification)
}
