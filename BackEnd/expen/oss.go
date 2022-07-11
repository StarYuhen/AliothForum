package expen

import (
	"BackEnd/config"
	"context"
	"github.com/sirupsen/logrus"
	"strings"
)

// 对接腾讯oss
// 推荐文章地址:https://www.shuzhiduo.com/A/1O5Eq94a57/

// TxOssBuckets 使用腾讯oss的结构体对象
type TxOssBuckets struct {
	FilePathClient  string `json:"FilePathClient,omitempty"`  // Buckets上传路径
	FileReader      string `json:"FileReader,omitempty"`      // 文件的字符串对象名
	FilePathBackEnd string `json:"FilePathBackEnd,omitempty"` // 本机地址
}

// TxUploadBuckets 向储存桶上传文件
func (t *TxOssBuckets) TxUploadBuckets() bool {
	// 上传文件夹对象,io流
	fi := strings.NewReader(t.FileReader)

	// put提交请求
	if _, err := config.TxOss.Object.Put(context.Background(), t.FilePathClient, fi, nil); err != nil {
		logrus.Info("put 预处理失败:", err)
		return false
	}
	// 上传本地文件
	if _, err := config.TxOss.Object.PutFromFile(context.Background(), t.FilePathClient, t.FilePathBackEnd, nil); err != nil {
		logrus.Info("上传文件失败:", err)
		return false
	}
	return true
}
