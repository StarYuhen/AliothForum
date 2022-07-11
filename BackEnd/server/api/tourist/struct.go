package tourist

// 存放api的struct的结构体文件

// LoginBoolTrue 登录接口返回参数
type LoginBoolTrue struct {
	ImgUrl string `json:"ImgUrl"`
	Jwt    string `json:"Jwt"`
	Name   string `json:"Name"`
}
