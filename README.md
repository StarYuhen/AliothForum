## 项目介绍

这是一个论坛网站，包含前后端

## 项目技术功能优点

* 可使用腾讯云OSS作为储存，减少服务器带宽压力
* 使用Elasticsearch作为论坛搜索服务，可承受高并发和优秀的搜索服务
* ~~使用RabbitMQ消息队列，可支撑高并发和限流熔断~~
* 使用Redis的内容缓存服务，有效降低Mysql的承受压力
* 使用协程池，能够在接口下支撑高并发的耗时及定时任务
* 拥有丰富的自定义选择，在BackEnd/config/config.yaml中

## 项目优点

* 采用Elasticsearch,~~RabbitMQ~~,Redis,协程池(ants) 在高并发的接口服务上承受能力极强
* 前端采用Vant和Element Plus,Tailwind CSS实现的快速开发界面，同时界面也更为美观
* 后端采用检测爬虫自动返回渲染的HTML文件，有利于文章的SEO优化，让网站拥有更多自然浏览
* 采用前后端分离结构，方便快捷开发，同时方便用户开发
* 拥有高度自定义的选项功能，能够更加快速的定义论坛网站
* 会自动根据配置内容来进行限流和并发支持

## 使用框架和第三方软件

### 前端

* Vant 手机UI框架
* Element Plus PC端UI框架
* Vue 网站基础框架
* github-markdown-css 美化文章markdown解析结果
* axios 作为请求接口的框架
* Tailwind CSS 作为快速开发前端美化CSS库

### 后端

* gorm 数据库操作框架
* gin Web服务接口路由框架
* Elasticsearch 用于论坛网站的搜索服务
* ~~RabbitMQ 消息队列框架~~
* Mysql用于储存用户账号和其余内容
* Redis 用于缓冲，减少对Mysql的消耗
* Lute 用于解析markdown内容返回html内容
* ants 充当协程池内容，服用协程节省资源
* tollbooth 作为HTTP接口的限流中间件功能

## 感谢提供

感谢Jetbrains提供的IDE支持
<a href="https://www.jetbrains.com/?from=go-wechaty"><img src="/docs/images/goland.png" width = "75px" height = "75px" alt="goland.png" /></a>



