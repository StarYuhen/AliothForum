package config

import (
	_ "embed"
	"github.com/88250/lute"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
	"gopkg.in/yaml.v2"
	"time"
)

type StructConfig struct {
	Version             version             `yaml:"Version"`
	MYSQLConfig         mysqlConfig         `yaml:"MYSQLConfig"`
	RedisConfig         redisConfig         `yaml:"RedisConfig"`
	WebEncrypt          webEncrypt          `yaml:"WebEncrypt"`
	WebFile             webFile             `yaml:"WebFile"`
	OSS                 oss                 `yaml:"OSS"`
	WebConfig           webConfig           `yaml:"WebConfig"`
	RabBitMQ            rabBitMQ            `yaml:"RabBitMQ"`
	ElasticsearchConfig elasticsearchConfig `yaml:"ElasticsearchConfig"`
}

type version struct {
	BackEnd        float64 `yaml:"BackEnd"`
	Web            float64 `yaml:"Web"`
	WebArticle     string  `yaml:"WebArticle"`
	BackEndArticle string  `yaml:"BackEndArticle"`
	GinMode        string  `yaml:"GinMode"`
}

type mysqlConfig struct {
	User         string   `yaml:"User"`
	Password     string   `yaml:"Password"`
	Architecture string   `yaml:"Architecture"`
	Addr         string   `yaml:"Addr"`
	Table        []string `yaml:"Table"`
}

type redisConfig struct {
	Addr         string `yaml:"Addr"`
	Password     string `yaml:"Password"`
	DB           []int  `yaml:"DB"`
	PoolSize     int    `yaml:"PoolSize"`
	MinIdleConns int    `yaml:"MinIdleConns"`
}

type webEncrypt struct {
	JWT        string `yaml:"JWT"`
	CookieURL  string `yaml:"CookieURL"`
	AESKey     string `yaml:"AESKey"`
	AESEncrypt string `yaml:"AESEncrypt"`
	WebURL     string `yaml:"WebURL"`
	Name       string `yaml:"Name"`
	QQToken    string `yaml:"QQToken"`
	QQMail     string `yaml:"QQMail"`
	RSAPrivate string `yaml:"RSAPrivate"`
	MD5Str     string `yaml:"MD5Str"`
}

type webFile struct {
	FileDir        string   `yaml:"FileDir"`
	Upload         string   `yaml:"Upload"`
	UploadImg      string   `yaml:"UploadImg"`
	UploadFile     string   `yaml:"UploadFile"`
	QrCodeURL      string   `yaml:"QrCodeURL"`
	MarkdownURL    string   `yaml:"MarkdownURL"`
	LogFile        string   `yaml:"LogFile"`
	UploadFileType []string `yaml:"UploadFileType"`
	UploadFileSize int      `yaml:"UploadFileSize"`
	UserImgUrl     string   `yaml:"UserImgUrl"`
}

type rabBitMQ struct {
	User        string `yaml:"User"`
	PassWord    string `yaml:"PassWord"`
	Addr        string `yaml:"Addr"`
	VirtualHost string `yaml:"VirtualHost"`
}

type oss struct {
	Bool     bool     `yaml:"Bool"`
	TxConfig txConfig `yaml:"TxConfig"`
}

type txConfig struct {
	SecretID  string `yaml:"SecretID"`
	SecretKey string `yaml:"SecretKey"`
	OSSUrl    string `yaml:"OSSUrl"`
}

type webConfig struct {
	MaxAllowed               int                      `yaml:"MaxAllowed"`
	GinPort                  string                   `yaml:"GinPort"`
	Log                      bool                     `yaml:"Log"`
	UTF8Number               int                      `yaml:"UTF8Number"`
	UserImgUrl               string                   `yaml:"UserImgUrl"`
	Ant                      int                      `yaml:"Ant"`
	NewRateLimiterMiddleware newRateLimiterMiddleware `yaml:"NewRateLimiterMiddleware"`
	TollBooth                tollBoothStruct          `yaml:"TollBooth"`
}

type newRateLimiterMiddleware struct {
	Key   string `yaml:"Key"`
	Limit int    `yaml:"Limit"`
	Time  int    `yaml:"Time"`
}

type tollBoothStruct struct {
	TokenBuckets         float64       `yaml:"TokenBuckets"`
	DefaultExpirationTTL time.Duration `yaml:"DefaultExpirationTTL"`
	Methods              []string      `yaml:"Methods"`
}

type elasticsearchConfig struct {
	URL      string `yaml:"URL"`
	User     string `yaml:"User"`
	PassWord string `yaml:"PassWord"`
	Gzip     bool   `yaml:"Gzip"`
}

//go:embed config.yaml
var FileConfig []byte

// StoreConfig 初始化配置内容和驱动器
var StoreConfig = Init()

// MysqlURL gorm的mysql数据库驱动
var MysqlURL = sqlInit()

// RedisWebExpen redis的缓存器驱动web杂项的
var RedisWebExpen = redisWebExpen()

// RedisArticle 储存文章内容
var RedisArticle = redisArticle()

// RedisComment 储存文章的点赞，收藏，ip
var RedisComment = redisComment()

// MailQQ QQ邮件发送驱动器
var MailQQ = gomail.NewDialer("smtp.qq.com", 587, StoreConfig.WebEncrypt.QQMail, StoreConfig.WebEncrypt.QQToken)

// Pool 协程池驱动器
var Pool = AntsPoolNew()

// TxOss 腾讯云对象储存驱动器
var TxOss = TxOSS()

// LuteEngine 启用MD文章解析框架 lute
var LuteEngine = lute.New()

// ElasticsearchEngine elasticsearch 搜索服务和文章储存服务
var ElasticsearchEngine = ElasticsearchClient()

// Limit 初始化速率限制器
var Limit = TooBooth()

// CacheAPI 设置api的接口缓存redis
var CacheAPI = persist.NewRedisStore(RedisWebExpen)

// RabBitMQEngine 启动消息队列
// 单机用个der MQ
// var RabBitMQEngine = RabBitMQ()

// Init 初始化配置文件内容
func Init() StructConfig {
	var con StructConfig
	// 将文件内容映射进入结构体
	if err := yaml.UnmarshalStrict(FileConfig, &con); err != nil {
		logrus.Error("读取文件映射结构体出错")
		panic(err)
	}
	return con
}

// DirInit 初始化文件存放目录
func DirInit() {
	MkDir(path+StoreConfig.WebFile.UploadImg, perm)
	MkDir(path+StoreConfig.WebFile.UploadFile, perm)
	MkDir(path+StoreConfig.WebFile.QrCodeURL, perm)
	MkDir(path+StoreConfig.WebFile.MarkdownURL, perm)
	MkDir(path+StoreConfig.WebFile.LogFile, perm)
	logrus.Info("初始化文件储存目录成功")
}
