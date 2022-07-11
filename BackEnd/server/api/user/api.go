package user

import (
	"BackEnd/config"
	"BackEnd/expen"
	"BackEnd/service"
	"BackEnd/service/UserAccountTable"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/skip2/go-qrcode"
	"net/http"
	"os"
)

// UploadImgUser 修改用户头像
func UploadImgUser(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	uid := ctx.MustGet("uid").(string)
	if err != nil {
		logrus.Error("上传文件错误--->", err)
		ctx.JSON(http.StatusOK, expen.MissingParametersFun("上传文件错误"))
		return
	}

	logrus.Info("传输的文件名字--->", file.Filename)
	suffix := PathText(file.Filename)
	if FileType(suffix) {
		logrus.Error("传输文件名字不对--->", err)
		ctx.JSON(http.StatusOK, expen.MissingParametersFun("传输文件后缀名错误"))
		return
	}
	// 设置保存的文件名
	FilePathNew := expen.CreatUid() + suffix
	// 设置本机oss路径
	ossPath := config.StoreConfig.WebFile.UploadImg + FilePathNew
	// 保存文件
	err = ctx.SaveUploadedFile(file, "./"+ossPath)

	if err != nil {
		logrus.Error("保存文件错误--->", err)
		ctx.JSON(http.StatusOK, expen.InternalErrorFun("保存文件错误"))
		return
	}
	dir, _ := os.Getwd()
	// 上传到oss
	var tx = new(expen.TxOssBuckets)
	tx.FileReader = FilePathNew
	tx.FilePathClient = ossPath
	tx.FilePathBackEnd = dir + "/" + ossPath
	if !tx.TxUploadBuckets() {
		logrus.Error("文件上传oss失败")
		ctx.JSON(http.StatusOK, expen.InternalErrorFun("上传oss错误"))
		return
	}

	// 修改数据库的img链接，并返回img给web网页使其动态更改头像地址
	img := config.StoreConfig.OSS.TxConfig.OSSUrl + "/" + ossPath
	if !UserAccountTable.UploadImg(uid, img) {
		ctx.JSON(http.StatusOK, expen.UnknownErrorFun("修改头像失败"))
		return
	}
	ctx.JSON(http.StatusOK, expen.Success(img, "修改头像成功"))
}

// JwtDevice 检查jwt包含的uid,并查询数据
func JwtDevice(ctx *gin.Context) {
	uid := ctx.MustGet("uid").(string)
	ctx.JSON(http.StatusOK, expen.Success(UserAccountTable.JwtDevice(uid), "查询uid成功"))
}

// ImgUser 根据jwt包含的uid查询头像地址
func ImgUser(ctx *gin.Context) {
	uid := ctx.MustGet("uid").(string)
	ctx.JSON(http.StatusOK, expen.Success(UserAccountTable.UserImg(uid), "查询头像地址成功"))
}

// UrlQrCode 根据链接生成二维码
func UrlQrCode(ctx *gin.Context) {
	// 先查询是否redis中缓存着有，如果有则返回已经生成的地址，没有则生成文件并分享地址，使用oss
	qr := ctx.Query("qrcode")
	// var tx = new(expen.TxOssBuckets)
	uid := uuid.NewString() + ".png"
	if Boole, url := expen.HaseRead(config.RedisWebExpen, "qrcode", qr); Boole {
		// 拥有缓存的二维码地址
		ctx.JSON(http.StatusOK, expen.Success(url, "请求二维码成功"))
		return
	}
	// ossPath := config.StoreConfig.WebFile.QrCodeURL + uid
	// 倘若未生成过则生成并上传oss
	if err := qrcode.WriteFile(qr, qrcode.Medium, 256, config.StoreConfig.WebFile.QrCodeURL+"/"+uid); err == nil {
		imgURL := config.StoreConfig.WebEncrypt.WebURL + "/QrCodeURL/" + uid
		// 插入redis
		expen.HaseSet(config.RedisWebExpen, "qrcode", qr, imgURL)
		ctx.JSON(http.StatusOK, expen.Success(imgURL, "生成二维码成功"))
		return
	}

	// 直接使用服务器自己的流量，不使用oss
	ctx.JSON(http.StatusOK, expen.InternalErrorFun("生成二维码失败"))

	// // 上传oss
	// tx.FileReader = uid
	// tx.FilePathClient = ossPath
	// tx.FilePathBackEnd = dir + "/" + ossPath
	// if Bool := tx.TxUploadBuckets(); Bool {
	// 	imgURL := config.StoreConfig.OSS.TxConfig.OSSUrl + "/" + ossPath
	// 	// 插入redis
	// 	expen.HaseSet("qrcode", qr, imgURL)
	// 	ctx.JSON(http.StatusOK, expen.Success(imgURL, "生成二维码成功"))
	// 	return
	// } else {
	// 	ctx.JSON(http.StatusOK, expen.InternalErrorFun("生成二维码失败"))
	// }

}

// InsertComment 插入文章评论
func InsertComment(ctx *gin.Context) {

}

// ReadComment 读取文章评论
func ReadComment(ctx *gin.Context) {
	var PostComment CommentRead
	if err := ctx.BindJSON(&PostComment); err != nil {
		logrus.Error("绑定JSON元素失败:", err)
		ctx.JSON(http.StatusOK, expen.ParameterErrorFun("请求数据错误，稍后再试"))
		return
	}
	index, max := 0, 0
	// 使用分页器获取内容
	if PostComment.Number == 1 {
		index = 1
		max = 10
	} else {
		index = (PostComment.Number-1)*10 + 1
		max = PostComment.Number * 10
	}
	logrus.Info("绑定的值：", PostComment)
	list, _ := service.PaginationCommentOne(index, max, PostComment.UID, PostComment.ClassificationUID+"_comment")
	ctx.JSON(http.StatusOK, expen.Success(list, "评论请求成功"))
}

// ClickLike 请求点赞
func ClickLike(ctx *gin.Context) {

}
