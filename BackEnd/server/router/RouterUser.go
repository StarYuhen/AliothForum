package router

import (
	"BackEnd/server/api/user"
	"github.com/gin-gonic/gin"
)

// UserRouter  用户登录后可以进行的操作
func UserRouter(engine *gin.RouterGroup) {
	// 上传图片并修改头像接口
	engine.POST("/UploadImg", user.UploadImgUser)
	// 检查jwt是否正确接口
	engine.GET("/JwtDevice", user.JwtDevice)
	// 查询用户的头像地址
	engine.GET("/ImgUser", user.ImgUser)
	// 根据链接生成二维码use.MajorIPSleepTwo(),
	engine.GET("/UrlQrCode", user.UrlQrCode)
	// 获取评论数据
	engine.POST("/ReadComment", user.ReadComment)
	// 增加点赞
	engine.GET("/AddLike", user.ClickLike)
	// 上传文件，用于贴子，评论，文章
	engine.POST("/ArticleUploadFile", user.ArticleUploadFile)
	// 文章点赞
	engine.GET("/ArticleLike", user.ArticleLike)
	// 所有收藏的贴吧
	engine.GET("/CollectedTab", user.CollectedTab)
	// 收藏该贴吧
	engine.GET("/CollectTab", user.CollectTab)
}
