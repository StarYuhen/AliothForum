package ForumListTable

import "BackEnd/config"

// ForumList 定义论坛总表数据模型
type ForumList struct {
	ID        uint   `gorm:"primaryKey"`
	UID       string `gorm:"column:uid"`
	Name      string `gorm:"column:Name"`
	ImgURL    string `gorm:"column:ImgUrl"`
	Src       string `gorm:"column:src"`
	CreateUID string `gorm:"column:createUid"`
}

// ForumIMGAndName 读取论坛头像和昵称
func ForumIMGAndName(uid string) ForumList {
	var f ForumList
	config.MysqlURL.Table("forum_list").Select("id", "uid", "Name", "ImgUrl", "Src").Where("uid", uid).First(&f)
	return f
}

// ForumListCrete 新建论坛信息
func (f *ForumList) ForumListCrete() error {
	return config.MysqlURL.Table("forum_list").Create(&f).Error
}

// RandomRecommend 随机推荐论坛-8个
func RandomRecommend() ([]ForumList, error) {
	var f []ForumList
	err := config.MysqlURL.Table("forum_list").
		Select("id", "Name", "ImgUrl", "Src").Where("id>=(?)",
		config.MysqlURL.Table("forum_list ").Select("FLOOR( MAX(id) * RAND())")).Limit(8).
		Find(&f).Error
	return f, err
}

// CollectedAll 获取所有已经关注了的论坛
func CollectedAll(all []string) ([]ForumList, error) {
	var f []ForumList
	err := config.MysqlURL.Table("forum_list").Where("uid in ?", all).Find(&f).Error
	return f, err
}
