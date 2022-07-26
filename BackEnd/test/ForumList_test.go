package test

import (
	"BackEnd/service/ForumListTable"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"testing"
)

// 创建论坛
func TestCreateList(t *testing.T) {
	var f ForumListTable.ForumList
	f.UID = uuid.NewString()
	f.CreateUID = "88f50918-cde0-11ec-94fb-8c8caad3310e"
	f.Name = "测试论坛"
	f.Src = "用于本地测试论坛"
	f.ImgURL = "https://q1.qlogo.cn/g?b=qq&nk=3446623843&s=640"
	logrus.Error(f.ForumListCrete())
}

// 获取随机推荐论坛
func TestRandomRecommend(t *testing.T) {
	logrus.Info(ForumListTable.RandomRecommend())
}
