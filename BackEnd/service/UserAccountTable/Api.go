package UserAccountTable

import (
	"BackEnd/config"
	"BackEnd/expen"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type RegisterAccountStruct struct {
	User     string `yaml:"User"`
	UserName string `yaml:"UserName"` // 用户的昵称
	Password string `yaml:"Password"`
	Captcha  string `yaml:"Captcha"` // 获取的captcha验证码的值
}

type LoginAccount struct {
	User     string `json:"User,omitempty"`
	Password string `json:"Password,omitempty"`
	Captcha  string `json:"Captcha,omitempty"`
}

type Register struct {
	ID       uint   `gorm:"primaryKey"`
	User     string `json:"User,omitempty"`
	UserName string `json:"UserName,omitempty" gorm:"column:username"`
	UID      string `json:"UID,omitempty" gorm:"column:uid"`
	Password string `json:"Password,omitempty" `
	Auto     string `json:"Auto,omitempty"`
	Times    int64  `json:"Times"`
	Img      string `json:"Img"`
}

// RegisterCreat 注册账号
func (r *Register) RegisterCreat(reg RegisterAccountStruct) bool {
	r.UserName = reg.UserName
	r.User = reg.User
	if FindAccount(r.User) {
		return false
	}
	r.Password = reg.Password
	r.Auto = "普通用户"
	r.Times = time.Now().Unix()
	r.UID = expen.CreatUid()
	r.Img = config.StoreConfig.WebFile.UserImgUrl
	// 是QQ邮箱则是QQ头像否则就是默认头像
	if QQMail := strings.Contains(r.User, "@qq.com"); QQMail {
		// 分割邮箱
		qq := strings.Split(r.User, "@")
		r.Img = "https://q1.qlogo.cn/g?b=qq&nk=" + qq[0] + "&s=640"
	}
	// 创建用户
	if err := config.MysqlURL.Table("user_account").Create(&r).Error; err != nil {
		logrus.Error("创建用户失败，原因:", err)
		return false
	}
	return true
}

// FindAccount 查询是否已注册账号
func FindAccount(u string) bool {
	var R Register
	config.MysqlURL.Table("user_account").Select("user").Where("user=?", u).First(&R)
	return R.User == u
}

// Login 登录账号
func (l *LoginAccount) Login() (bool, *Register) {
	var R *Register
	config.MysqlURL.Table("user_account").Where("user=? and password=?", l.User, l.Password).First(&R)
	return R.User == l.User, R
}

// UploadImg 修改头像地址
func UploadImg(uid, img string) bool {
	if err := config.MysqlURL.Table("user_account").Where("uid=?", uid).Update("img", img).Error; err != nil {
		logrus.Error("更新用户头像地址失败，原因:", err)
		return false
	}
	return true
}

// JwtDevice  检查jwt是否正确
func JwtDevice(uid string) bool {
	var r Register
	config.MysqlURL.Table("user_account").Select("uid").Where("uid", uid).First(&r)
	return r.UID == uid
}

// UserIMGAndName 获取数据库中的头像和昵称
func UserIMGAndName(uid string) Register {
	var r Register
	config.MysqlURL.Table("user_account").Select("username", "img").Where("uid", uid).First(&r)
	return r
}

// UserImg 根据uid返回img头像地址
func UserImg(uid string) string {
	var r Register
	config.MysqlURL.Table("user_account").Where("uid", uid).First(&r)
	return r.Img
}
