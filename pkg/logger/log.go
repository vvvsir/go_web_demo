package logger

import (
	"go_web_demo/pkg/setting"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func Setup() {
	//写入文件
	if setting.LogConf.Path == "" {
		setting.LogConf.Path, _ = os.Getwd()
		setting.LogConf.Path += "/log"
	}

	level, err := logrus.ParseLevel(setting.LogConf.Level)
	if err != nil {
		log.Panic("日志level格式设置错误", err)
	}
	//设置最低loglevel
	Logger.SetLevel(level)

	//设置日志格式
	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	Logger.SetFormatter(&logrus.JSONFormatter{})
	// Logger.SetFormatter(&logrus.TextFormatter{}) // 默认格式 无需显示设置

	// 取消线程安全
	Logger.SetNoLock()

	// 自定义HOOK
	Logger.AddHook(&GHook{})
}
