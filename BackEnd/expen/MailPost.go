package expen

import (
	"BackEnd/config"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

// MailPOSTNumber 发送验证码
func MailPOSTNumber(To string, Number string) bool {
	m := gomail.NewMessage()
	// 发送人
	m.SetHeader("From", config.StoreConfig.WebEncrypt.QQMail)
	// 接收人
	m.SetHeader("To", To)
	// 主题
	m.SetHeader("Subject", "Alioth论坛验证码")
	// 内容
	body := fmt.Sprintf("<a href=%s style=color:#12addb target=_blank>%s </a> <h2>请求注册账号验证码：%s<h2> <p>此验证码两分钟后失效。</p>  <p>若此邮件不是您请求的，请忽略并删除！</p> ",
		config.StoreConfig.WebEncrypt.WebURL,
		config.StoreConfig.WebEncrypt.Name,
		Number)
	m.SetBody("text/html", body)

	// 发送邮件
	if err := config.MailQQ.DialAndSend(m); err != nil {
		logrus.Error("发送邮件失败，原因:", err)
		return false
	}
	logrus.Info("发送邮件成功")
	return true
}

// MailNoticeRegisterAccount 发送注册账号成功通知
func MailNoticeRegisterAccount(User, Username string) bool {
	m := gomail.NewMessage()
	// 发送人
	m.SetHeader("From", config.StoreConfig.WebEncrypt.QQMail)
	// 接收人
	m.SetHeader("To", User)
	// 主题
	m.SetHeader("Subject", config.StoreConfig.WebEncrypt.Name+"注册成功邮件")
	// 内容
	body := fmt.Sprintf("<a href=%s style=color:#12addb target=_blank> %s </a> <p>恭喜您在%s 注册成功！</p> <p>登录账号为：%s</p> <p>昵称为:%s </p>   <p>若此邮件不是您请求的，请忽略并删除！</p> ",
		config.StoreConfig.WebEncrypt.WebURL,
		config.StoreConfig.WebEncrypt.Name,
		config.StoreConfig.WebEncrypt.Name,
		User,
		Username)
	logrus.Info(body)

	m.SetBody("text/html", body)
	// 发送邮件
	if err := config.MailQQ.DialAndSend(m); err != nil {
		logrus.Error("发送邮件失败，原因:", err)
		return false
	}
	logrus.Info("发送邮件成功")
	return true
}

// MailLoginAccount 发送登录成功邮件通知
func MailLoginAccount(User, Username, IP string) bool {
	m := gomail.NewMessage()
	// 发送人
	m.SetHeader("From", config.StoreConfig.WebEncrypt.QQMail)
	// 接收人
	m.SetHeader("To", User)
	// 主题
	m.SetHeader("Subject", config.StoreConfig.WebEncrypt.Name+"登录成功邮件")
	// 内容
	body := fmt.Sprintf("<a href=%s style=color:#12addb target=_blank> %s </a> <p>恭喜您在%s 登录成功！</p> <p>登录账号为：%s</p> <p>昵称为:%s </p>  <p>登录IP:%s </p> <p>若此邮件不是您请求的，请忽略并删除！</p> ",
		config.StoreConfig.WebEncrypt.WebURL,
		config.StoreConfig.WebEncrypt.Name,
		config.StoreConfig.WebEncrypt.Name,
		User,
		Username,
		IP)
	logrus.Info(body)

	m.SetBody("text/html", body)
	// 发送邮件
	if err := config.MailQQ.DialAndSend(m); err != nil {
		logrus.Error("发送邮件失败，原因:", err)
		return false
	}
	logrus.Info("发送邮件成功")
	return true
}
