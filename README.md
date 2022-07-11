## 项目介绍

这是一个论坛网站，包含前后端

## 感谢提供

感谢Jetbrains提供的IDE支持

## 项目技术功能优点

* 可使用腾讯云OSS作为储存，减少服务器带宽压力
* 使用Elasticsearch作为论坛搜索服务，可承受高并发和优秀的搜索服务
* 使用RabbitMQ消息队列，可支撑高并发和限流熔断
* 使用Redis的内容缓存服务，有效降低Mysql的承受压力
* 使用协程池，能够在接口下支撑高并发的耗时及定时任务
* 拥有丰富的自定义选择，在Config/config.yaml 中

## 使用框架和第三方软件

### 前端

* Vant 手机UI框架
* Element Plus PC端UI框架
* Vue 网站基础框架
* github-markdown-css 美化文章markdown解析结果
* axios 作为请求接口的框架

### 后端

* gorm 数据库操作框架
* gin Web服务接口路由框架
* Elasticsearch 用于论坛网站的搜索服务
* RabbitMQ 消息队列框架
* Mysql用于储存用户账号和其余内容
* Redis 用于缓冲，减少对Mysql的消耗
* Lute 用于解析markdown内容返回html内容
* ants 充当协程池内容，服用协程节省资源
* tollbooth 作为HTTP接口的限流中间件功能




