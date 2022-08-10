package tourist

import (
	"BackEnd/config"
	"BackEnd/expen"
	"BackEnd/server/api/function"
	"BackEnd/server/use"
	"BackEnd/service"
	"BackEnd/service/ForumListTable"
	"BackEnd/service/UserAccountTable"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Captcha 图型验证码接口
func Captcha(ctx *gin.Context) {
	id, b64, number := UserAccountTable.NumberCaptcha() // 生成验证码
	// 写入redis
	expen.StringSet(config.RedisWebExpen, id, number, time.Second*60)
	// 设置cookie
	ctx.SetCookie("CaptchaCache", id, 60, "/", config.StoreConfig.WebEncrypt.CookieURL, false, true)
	// 返回图标
	ctx.JSON(http.StatusOK, expen.Success(b64, "请求图像验证码成功"))
}

// RegisterAccount 注册论坛账号
func RegisterAccount(ctx *gin.Context) {
	var reg UserAccountTable.RegisterAccountStruct
	var r = new(UserAccountTable.Register)
	CaptchaID, _ := ctx.Cookie("CaptchaCache")
	if err := ctx.BindJSON(&reg); err != nil {
		logrus.Error("绑定JSON元素失败:", err)
		ctx.JSON(http.StatusOK, expen.ParameterErrorFun("注册账号失败"))
		return
	}

	// 校检验证码
	boole, err := expen.StringRead(config.RedisWebExpen, CaptchaID)
	ma := strings.Split(err, "|")
	if boole && ma[1] == reg.Captcha && ma[0] == reg.User {
		// 检验成功后注册账号
		if booles := r.RegisterCreat(reg); booles {
			// 创建用户成功后设置cookie
			jwt, _ := use.GenToken(r.User, r.UID, r.Auto)

			// 发送注册成功邮件
			config.Pool.Submit(func() {
				expen.MailNoticeRegisterAccount(r.User, r.UserName)
			})

			ctx.JSON(http.StatusOK, expen.Success(jwt, "注册账号成功"))
		} else {
			ctx.JSON(http.StatusOK, expen.InternalErrorFun("创建账号失败,可能账号已存在"))
		}

		return
	}

	ctx.JSON(http.StatusOK, expen.ParameterErrorFun("提交的验证码错误！"))
}

// LoginAccount 登录账号
func LoginAccount(ctx *gin.Context) {
	var login UserAccountTable.LoginAccount
	CaptchaID, _ := ctx.Cookie("CaptchaCache")
	if err := ctx.BindJSON(&login); err != nil {
		logrus.Error("绑定JSON元素失败:", err)
		ctx.JSON(http.StatusOK, expen.ParameterErrorFun("请求数据错误，稍后再试"))
		return
	}
	// 校检验证码
	// if 判断太多了，下次多条件使用面向对象写法
	if boole, err := expen.StringRead(config.RedisWebExpen, CaptchaID); boole && err == login.Captcha {
		// 检验成功后登录账号
		if booles, account := login.Login(); booles {
			// 创建用户成功后设置jwt
			var True LoginBoolTrue
			jwt, _ := use.GenToken(account.User, account.UID, account.Auto)

			// 发送注册成功邮件
			config.Pool.Submit(func() {
				expen.MailLoginAccount(account.User, account.UserName, ctx.ClientIP())
			})
			True.Name = account.UserName
			True.Jwt = jwt
			True.ImgUrl = account.Img

			ctx.JSON(http.StatusOK, expen.Success(True, "登录账号成功"))
		} else {

			if UserAccountTable.FindAccount(login.User) {
				ctx.JSON(http.StatusOK, expen.InternalErrorFun("登录账号失败,账号密码错误"))
			} else {
				ctx.JSON(http.StatusOK, expen.ParameterErrorFun("账号不存在"))
			}

		}

		return

	}

	ctx.JSON(http.StatusOK, expen.ParameterErrorFun("提交的验证码错误！"))

}

// RegisterAccountMail 发送注册邮件验证码
func RegisterAccountMail(ctx *gin.Context) {
	mail := ctx.Query("mail")
	id, _, number := UserAccountTable.NumberCaptcha() // 生成验证码
	// 写入redis
	expen.StringSet(config.RedisWebExpen, id, mail+"|"+number, time.Second*60)
	// 设置cookie
	ctx.SetCookie("CaptchaCache", id, 60, "/", config.StoreConfig.WebEncrypt.CookieURL, false, true)
	// 发送验证码
	if !expen.MailPOSTNumber(mail, number) {
		ctx.JSON(http.StatusOK, expen.InternalErrorFun("请求邮箱验证码失败"))
		return
	}
	ctx.JSON(http.StatusOK, expen.Success(nil, "请求邮箱验证码成功"))
}

// GetArticleContentReptile 判断爬虫的内容判定
func GetArticleContentReptile(ctx *gin.Context) {
	Article := new(function.ArticleContent)
	// 利用重定向，是浏览器则输出，不是浏览器则直接重定向
	// https://juejin.cn/post/6995911587412344840
	Article.ID = ctx.Param("ID")
	// 获取储存在文章redis的浏览量，点赞量，作者ID
	if err := Article.GetVal(); err != nil {
		logrus.Error("请求文章内容接口失败:", err)
		ctx.JSON(http.StatusOK, expen.InternalErrorFun("访问文章失败"))
		return
	}
	// 读取成功后默认增加浏览量
	Article.AddPageViews(ctx.ClientIP())
	// 输出html 代码 https://segmentfault.com/q/1010000012076052
	if value, Bool := ctx.Get("Reptile"); value.(bool) && Bool {
		// 都为真则是爬虫直接输出html内容
		ctx.Header("Content-Type", "text/html; charset=utf-8")
		html := config.LuteEngine.MarkdownStr(Article.ArticleData.Title, Article.Article.Content)
		// 用于模板渲染返回html title需要接上网站标题
		// 随机网址
		ctx.HTML(http.StatusOK, "Article.tmpl", gin.H{
			"Title":          Article.ArticleData.Title + "-" + config.StoreConfig.WebEncrypt.Name,
			"description":    Article.Article.Description,
			"TitleContent":   Article.ArticleData.Title,
			"keywords":       Article.Article.Keywords,
			"ArticleContent": template.HTML(html),
			"RandArticleURL": config.StoreConfig.Version.BackEndArticle + expen.HashRandomKey(config.RedisArticle),
		})
	} else {
		// 重定向到前端文章地址
		ctx.Redirect(http.StatusMovedPermanently, config.StoreConfig.Version.WebArticle+Article.ID)
	}
}

// GetArticleContent 获取指定储存的网页论坛内容信息
/*
	无论是文章还是评论统一写入redis储存，mysql慢查询太容易触发了
	使用vditor 解析md和编写md https://github.com/Vanessa219/vditor 前端
	后端使用 https://github.com/88250/lute 作为md解析 --暂时不采用后端作为md解析
	文章储存逻辑，编写内容后上传完整的md文件，接口返回md路径和其他参数
	其他参数：浏览量，点赞量，作者ID,评论各项参数
*/
func GetArticleContent(ctx *gin.Context) {
	Article := new(function.ArticleContent)
	Article.ID = ctx.Param("ID")
	// 获取储存在文章redis的浏览量，点赞量，作者ID
	if err := Article.GetVal(); err != nil {
		logrus.Error("请求文章内容接口失败:", err)
		ctx.JSON(http.StatusOK, expen.InternalErrorFun("访问文章失败"))
		return
	}
	// 读取成功后默认增加浏览量
	Article.AddPageViews(ctx.ClientIP())
	ctx.JSON(http.StatusOK, expen.Success(Article, "欢迎查看文章"))
}

// RandomRecommendForum 随机获取推荐论坛
// 不需要判断是否收藏过
func RandomRecommendForum(ctx *gin.Context) {
	list, err := ForumListTable.RandomRecommend()
	if err != nil {
		logrus.Error("获取随机推荐论坛失败:", err)
		ctx.JSON(http.StatusOK, expen.InternalErrorFun("获取随机推荐论坛失败"))
		return
	}
	ctx.JSON(http.StatusOK, expen.Success(list, strconv.Itoa(len(list))))
}

// ArticleRandomIO 获取随机推荐文章
func ArticleRandomIO(ctx *gin.Context) {
	// 从redis读取文章ID，然后查询数据并行获取结果
	List := make([]struct {
		All     *function.ArticleRedis `json:"All"`
		Article *service.Table         `json:"Article"`
	}, 7)
	for i := 0; i <= 6; i++ {
		id := expen.HashRandomKey(config.RedisArticle)
		// 通过key查找value
		str := expen.HashReadAll(config.RedisArticle, id)
		// 直接遍历到redis 结构体内容去
		all := new(function.ArticleRedis)
		sql := new(service.Table)
		if err := mapstructure.WeakDecode(str, all); err != nil {
			logrus.Error("转换Redis中的文章数据失败:", err)
			continue
		} else {
			// 再转换一边json
			if err := json.Unmarshal([]byte(all.Mysql), sql); err != nil {
				logrus.Error("转换json中的文章数据失败:", err)
				continue
			}
			List[i].Article = sql
			List[i].All = all
		}

	}
	// 直接返回结果
	ctx.JSON(http.StatusOK, expen.Success(List, "请求首页随机文章成功"))

}

// ForumDetails 获取指定论坛信息
func ForumDetails(ctx *gin.Context) {
	UID := ctx.Query("uid")
	f := ForumListTable.ForumIMGAndName(UID)
	ctx.JSON(http.StatusOK, expen.Success(f, "请求论坛信息成功"))
}

// ForumArticle 获取论坛内的随机推荐内容
func ForumArticle(ctx *gin.Context) {
	// 获取数据库的随机推荐文章
	uid := ctx.Query("uid")
	// 从redis读取文章ID，然后查询数据并行获取结果
	List := make([]struct {
		All     *function.ArticleRedis `json:"All"`
		Article *service.Table         `json:"Article"`
	}, 6)
	t, err := service.RandomArticle(uid)

	if err != nil {
		logrus.Error("请求随机推荐文章失败:", err)
		ctx.JSON(http.StatusOK, expen.HttpToast(err, "请求随机推荐文章失败"))
		return
	}

	for i, v := range t {
		// 查询的redis数据插入进去
		str := expen.HashReadAll(config.RedisArticle, v.Uid)
		// 直接遍历到redis 结构体内容去
		all := new(function.ArticleRedis)
		if err := mapstructure.WeakDecode(str, all); err != nil || all == nil {
			logrus.Error("转换Redis中的文章数据失败:", err)
			continue
		} else {
			List[i].All = all
		}
		List[i].Article = &v
	}
	ctx.JSON(http.StatusOK, expen.Success(List, "请求随机推荐文章成功"))
}
