package expen

import (
	"github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

// StringsP 匹配p标签内容
func StringsP(html string) string {
	var hrefRegexp = regexp.MustCompile("<p.*?>.*?</p>")
	match := hrefRegexp.FindAllString(html, -1)
	var content string
	if match == nil {
		return ""
	}
	for _, v := range match {
		// p标签提取
		list := strings.Split(strings.Split(v, ">")[1], "<")[0]
		content = content + list
		if len(content) > 550 {
			// 截取前50个
			logrus.Info(content)
			logrus.Info(content[:550])
			return content
		}
	}
	return content
}
