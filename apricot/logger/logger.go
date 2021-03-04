package logger

import (
	"io"
	"os"

	"github.com/alex-techs/apricot/apricot/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger = logrus.New()

func InitLog() {
	// 设置为 true 则日志将会输出颜色，并将内容解析为直观的样式进行输出
	Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: config.C.Log.ForceColors,
	})

	Logger.SetOutput(io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   config.C.Log.File,
		MaxSize:    config.C.Log.MaxAge,
		MaxBackups: config.C.Log.MaxBackup,
		MaxAge:     config.C.Log.MaxAge,
		Compress:   config.C.Log.Compress,
	}))
}
