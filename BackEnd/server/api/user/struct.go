package user

import "BackEnd/config"

// CommentRead 读取评论数据内容
type CommentRead struct {
	UID               string `json:"UID" binding:"required"`               // 文章uid
	Number            int    `json:"Number" `                              // 分页数据，1则是1-10，2则是11-20--谁用接口传范围谁是SB
	Type              bool   `json:"Type" `                                // 是一级还是二级
	CommentUID        string `json:"CommentUID"`                           // 是二级才会启用这个字段
	ClassificationUID string `json:"ClassificationUID" binding:"required"` // 属于哪个论坛
}

// // CommentCreate 创建评论数据
// type CommentCreate struct {
//
// }

type UploadImg struct {
}

// PathText 获取文件的扩展名
func PathText(path string) string {
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[i:]
		}
	}
	return ""
}

// FileType 判断文件后缀名是否符合规定
func FileType(Type string) bool {
	for _, v := range config.StoreConfig.WebFile.UploadFileType {
		if Type == v {
			return true
		}
	}
	return false
}
