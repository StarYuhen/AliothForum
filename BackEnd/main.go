package main

import (
	"BackEnd/config"
	"BackEnd/server"
	Rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"runtime"
	"time"
)

func main() {
	// 初始化文件目录
	config.DirInit()
	// 设置日志库增加代码信息排查
	logrus.SetReportCaller(true)
	// 设置日志库json格式输出日志
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// 利用日志分割工具设置每天分割一次日志 --配置文件修改
	if config.StoreConfig.WebConfig.Log {
		LogPath := config.StoreConfig.WebFile.LogFile + "/log"
		writerLog, _ := Rotatelogs.New(
			LogPath+".%Y%m%d%H%M",
			Rotatelogs.WithLinkName(LogPath),
			Rotatelogs.WithMaxAge(time.Duration(168)*time.Hour),        // 设置一个星期自动清理一次过去文件
			Rotatelogs.WithRotationTime(time.Duration(24)*time.Second), // 设置一天日志进行分割储存
		)
		logrus.SetOutput(writerLog)
	}
	// 设置协程抢占cpu核心数为最大
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 启动gin引擎
	server.GinServer()
	logrus.Info("已成功启动AliothForum 论坛服务！")
}
