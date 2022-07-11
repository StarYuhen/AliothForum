package major

import (
	"BackEnd/expen"
	"BackEnd/server/api/function"
	"github.com/gin-gonic/gin"
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
