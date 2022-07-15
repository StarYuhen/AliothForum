package major

import (
	"BackEnd/config"
	"BackEnd/expen"
	"BackEnd/server/api/function"
	"BackEnd/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

// InsertArticle 新建文章
func InsertArticle(ctx *gin.Context) {
	var Article = new(function.ArticleContent)
	var Post ArticleContentCrateRequest
	if err := ctx.BindJSON(&Post); err != nil {
		logrus.Error("绑定JSON元素失败:", err)
		ctx.JSON(http.StatusOK, expen.ParameterErrorFun("请求数据错误，稍后再试"))
		return
	}

	Article.Article.Img = Post.Img
	Article.ArticleData.Title = Post.Title
	Article.ArticleData.ViewPermissions = Post.ViewPermissions
	Article.ArticleData.AuthorUID = Post.AuthorUID
	Article.ArticleData.ClassificationUID = Post.ClassificationUID
	Article.Article.Content = Post.Content
	Article.Article.Keywords = Post.Keywords

	Article.ID = strconv.FormatInt(time.Now().UnixNano(), 10)

	if Article.Init() {
		ctx.JSON(http.StatusOK, expen.Success(Article.ID, "创建文章成功"))
		return
	}
	ctx.JSON(http.StatusOK, expen.UnknownErrorFun("创建文章失败"))
}

// InsertComment 插入文章评论
func InsertComment(ctx *gin.Context) {
	var comment CommentCreate
	var commentTable service.Comment
	if err := ctx.BindJSON(&comment); err != nil {
		logrus.Error("绑定JSON元素失败:", err)
		ctx.JSON(http.StatusOK, expen.ParameterErrorFun("请求数据错误，稍后再试"))
		return
	}

	commentTable.Uid = comment.ArticleUID
	commentTable.AuthorName = comment.AuthorName
	commentTable.AuthImg = comment.AuthorIMG
	// 当选择是一级评论时
	if comment.Type {
		commentTable.CommentOneUid = uuid.NewString()
	} else {
		// 当选择是二级评论时
		commentTable.CommentOneUid = comment.CommentUID
	}
	commentTable.CommentText = config.LuteEngine.MarkdownStr(comment.ArticleUID, comment.Text)

	if service.CreateComment(commentTable, comment.ClassificationUID+"_comment") != nil {
		// 创评论数据失败
		ctx.JSON(http.StatusOK, expen.InternalErrorFun("评论失败"))
		return
	}
	// 创建成功
	ctx.JSON(http.StatusOK, expen.Success(commentTable.CommentOneUid, "创建评论数据成功"))
}

// CreateClassification 新建论坛--论坛用于储存论坛表
func CreateClassification(ctx *gin.Context) {
	// 创建论坛后依次创建论坛文章表和论坛评论表

}
