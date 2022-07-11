package major

// ArticleContentCrateRequest 创建文章接口内容
type ArticleContentCrateRequest struct {
	Title             string `json:"Title"`             // 文章标题
	Img               string `yaml:"Img"`               // 倘若有这是第一个上传的图像内容
	AuthorUID         string `json:"AuthorUID"`         // 文章归属作者uid
	ClassificationUID string `json:"ClassificationUID"` // 文章归属子论坛uid
	Content           string `json:"Content"`           // 文章内容
	ViewPermissions   bool   `json:"ViewPermissions"`   // 是否需要登录查看
	Keywords          string `json:"Keywords"`          // 倘若发送的时文章请自动增加关键词
}
