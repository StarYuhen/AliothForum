package major

// ArticleContentCrateRequest 创建文章接口内容
type ArticleContentCrateRequest struct {
	Title             string `json:"Title" binding:"required"`             // 文章标题
	Img               string `yaml:"Img" binding:"required"`               // 倘若有这是第一个上传的图像内容
	AuthorUID         string `json:"AuthorUID" binding:"required"`         // 文章归属作者uid
	ClassificationUID string `json:"ClassificationUID" binding:"required"` // 文章归属子论坛uid
	Content           string `json:"Content" binding:"required"`           // 文章内容
	ViewPermissions   bool   `json:"ViewPermissions"`                      // 是否需要登录查看
	Keywords          string `json:"Keywords" binding:"required"`          // 倘若发送的时文章请自动增加关键词
}

// CommentCreate 创建评论数据
type CommentCreate struct {
	ArticleUID        string `json:"ArticleUID" binding:"required"`        // 文章uid
	ClassificationUID string `json:"ClassificationUID" binding:"required"` // 用于读取时哪个分类论坛的信息
	Type              bool   `json:"Type" `                                // 插入评论，true是一级评论。false是二级评论
	CommentUID        string `json:"CommentUID"`                           // 倘若是二级评论请加上一级评论UID，倘若是一级评论则为空
	AuthorName        string `json:"AuthorName" binding:"required"`        // 评论者昵称
	AuthorIMG         string `json:"AuthorIMG" binding:"required"`         // 评论着头像
	Text              string `json:"Text" binding:"required"`              // 评论内容
}

// CreateForumPost 创建论坛请求数据
type CreateForumPost struct {
	Name   string `json:"Name" binding:"required"`
	Src    string `json:"Src" binding:"required"`
	ImgURL string `json:"ImgURL" binding:"required"`
}
