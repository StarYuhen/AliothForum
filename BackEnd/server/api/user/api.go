package user

import (
	"BackEnd/config"
	"BackEnd/expen"
	"BackEnd/service"
	"BackEnd/service/ForumListTable"
	"BackEnd/service/UserAccountTable"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/skip2/go-qrcode"
	"net/http"
	"time"
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
	FilePathNew := uuid.NewString() + suffix
	// 设置本机oss路径
	ossPath := config.StoreConfig.WebFile.UploadImg + FilePathNew
	// // 保存文件
	// err = ctx.SaveUploadedFile(file, "./"+ossPath)
	//
	// if err != nil {
	// 	logrus.Error("保存文件错误--->", err)
	// 	ctx.JSON(http.StatusOK, expen.InternalErrorFun("保存文件错误"))
	// 	return
	// }
	// 上传到oss
	var tx = new(expen.TxOssBuckets)
	tx.FileReader = FilePathNew
	tx.FilePathClient = ossPath
	tx.FilePathBackEnd, _ = file.Open()

	if tx.TxUploadBuckets() {
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
		index = 0
		max = 6
	} else {
		index = (PostComment.Number-1)*6 + 1
		max = PostComment.Number * 6
	}
	logrus.Info("绑定的值：", PostComment)
	list, _ := service.PaginationCommentOne(index, max, PostComment.UID, PostComment.ClassificationUID+"_comment")
	ctx.JSON(http.StatusOK, expen.Success(list, "评论请求成功"))
}

// ClickLike 请求点赞
func ClickLike(ctx *gin.Context) {
	// 查询文章Ip里面是否有相同的IP，单个文章相同IP不作数
	uid, _ := ctx.Get("uid")
	ID, _ := ctx.GetQuery("ID")
	if Null, _ := expen.HaseRead(config.RedisComment, ID, uid.(string)+"_like"); !Null {
		// 插入点赞UID和点赞数量
		expen.HaseSet(config.RedisComment, ID, uid.(string)+"_like", true)
		expen.HashInsertAdd(config.RedisArticle, ID, "Likes")
		ctx.JSON(http.StatusOK, expen.Success(nil, "点赞成功"))
	}
	ctx.JSON(http.StatusOK, expen.Success(nil, "您已点过赞"))
}

// ArticleUploadFile 上传文件
func ArticleUploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		logrus.Error("上传文件错误--->", err)
		ctx.JSON(http.StatusOK, expen.MissingParametersFun("上传文件错误"))
		return
	}
	logrus.Info("传输的文件名字--->", file.Filename)
	suffix := PathText(file.Filename)
	if !FileType(suffix) {
		logrus.Error("传输文件名字不对--->", err)
		ctx.JSON(http.StatusOK, expen.MissingParametersFun("传输文件后缀名错误"))
		return
	}
	// 设置保存的文件名
	FilePathNew := uuid.NewString() + suffix
	// 设置本机oss路径
	ossPath := config.StoreConfig.WebFile.UploadFile + FilePathNew
	// // 保存文件
	// err = ctx.SaveUploadedFile(file, "./"+ossPath)
	//
	// if err != nil {
	// 	logrus.Error("保存文件错误--->", err)
	// 	ctx.JSON(http.StatusOK, expen.InternalErrorFun("保存文件错误"))
	// 	return
	// }
	// 上传到oss
	var tx = new(expen.TxOssBuckets)
	tx.FileReader = FilePathNew
	tx.FilePathClient = ossPath
	tx.FilePathBackEnd, _ = file.Open()
	// 进行延时关闭
	defer tx.FilePathBackEnd.Close()
	if tx.TxUploadBuckets() {
		logrus.Error("文件上传oss失败:", err)
		ctx.JSON(http.StatusOK, expen.InternalErrorFun("上传oss错误"))
		return
	}
	img := config.StoreConfig.OSS.TxConfig.OSSUrl + "/" + ossPath
	ctx.JSON(http.StatusOK, expen.Success(img, "上传文件成功"))
	return
}

// ArticleLike 文章点赞
func ArticleLike(ctx *gin.Context) {
	// 判断redis是否有用户的点赞记录，有则提示点赞
	UserUID, _ := ctx.Get("uid")
	ArticleUID := ctx.Query("article")
	if bo, _ := expen.HaseRead(config.RedisComment, ArticleUID, UserUID.(string)); bo {
		//  为真说明已经点过赞了
		ctx.JSON(http.StatusOK, expen.InternalErrorFun("您已经点过赞了"))
		return
	}

	// 没有点过
	expen.HaseSet(config.RedisComment, ArticleUID, UserUID.(string), time.Now())
	expen.HashInsertAdd(config.RedisArticle, ArticleUID, "Likes")
	ctx.JSON(http.StatusOK, expen.Success(nil, "点赞成功惹"))
}

// CollectTab 收藏贴吧
func CollectTab(ctx *gin.Context) {
	UserUID, _ := ctx.Get("uid")
	TabUID := ctx.Query("tab")
	if b, _ := expen.HaseRead(config.RedisWebExpen, UserUID.(string), TabUID); b {
		ctx.JSON(http.StatusOK, expen.HttpToast(b, "您已经收藏了该贴吧"))
		return
	}
	// 没有收藏则收藏
	ctx.JSON(http.StatusOK,
		expen.Success(expen.HaseSet(config.RedisWebExpen, UserUID.(string), TabUID, true),
			"您已收藏成功"))
}

// CollectedTab 已经收藏了的贴吧
func CollectedTab(ctx *gin.Context) {
	// 收藏了的贴吧放进redis中
	UserUID, _ := ctx.Get("uid")
	all := expen.HashReadAllKey(config.RedisWebExpen, UserUID.(string))
	if len(all) == 0 {
		ctx.JSON(http.StatusOK, expen.HttpToast(all, "您暂未收藏贴吧"))
		return
	}
	// 通过redis返回的所有key查询数据库内的贴吧信息
	f, err := ForumListTable.CollectedAll(all)
	if err != nil {
		ctx.JSON(http.StatusOK, expen.MissingParametersFun("获取收藏贴吧数据失败"))
		return
	}
	ctx.JSON(http.StatusOK, expen.Success(f, "请求收藏贴吧成功"))
}
