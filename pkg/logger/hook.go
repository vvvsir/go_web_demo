package logger

import (
	"fmt"
	"go_web_demo/pkg/setting"
	"log"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type GHook struct{}

var openFile *os.File

func (h *GHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *GHook) Fire(entry *logrus.Entry) (err error) {
	fileName := time.Now().Format("2006-01-02")
	fullPath := fmt.Sprintf("%s/%s.log", setting.LogConf.Path, fileName)

	// 无需多次获取文件句柄
	if openFile != nil && openFile.Name() == fullPath {
		return
	}

	if err = os.MkdirAll(setting.LogConf.Path, os.ModePerm); err != nil {
		log.Panic("创建文件夹错误", err)
		return
	}

	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	openFile, err = os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Panic("写入日志文件错误", err)
		return
	}

	// //设置输出(只写文件)
	// Logger.Out = openFile

	/* 日志轮转相关函数
	`WithLinkName` 为最新的日志建立软连接
	`WithRotationTime` 设置日志分割的时间，隔多久分割一次
	 WithMaxAge 和 WithRotationCount二者只能设置一个
	`WithMaxAge` 设置文件清理前的最长保存时间
	`WithRotationCount` 设置文件清理前最多保存的个数
	*/
	// 下面配置日志每隔 RorateTime 小时轮转一个新文件，保留最近 MaxAge 小时的日志文件，多余的自动清理掉。
	writer, _ := rotatelogs.New(
		fullPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(fullPath),
		rotatelogs.WithMaxAge(time.Duration(setting.LogConf.MaxAge)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(setting.LogConf.RorateTime)*time.Hour),
	)
	Logger.Out = writer

	// //同时写文件和屏幕
	// writers := []io.Writer{openFile, os.Stdout}
	// fileAndStdoutWriter := io.MultiWriter(writers...)
	// Logger.Out = fileAndStdoutWriter

	return
}
