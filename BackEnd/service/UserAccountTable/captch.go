package UserAccountTable

import (
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

var store = base64Captcha.DefaultMemStore

func NumberCaptcha() (string, string, string) {
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		logrus.Info("验证码获取失败", err)
		return "", "", ""
	} else {
		number := cp.Store.Get(id, true)
		logrus.Printf("验证码ID:%s,对应数字:%s", id, number)
		return id, b64s, number
	}
}
