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
	Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: config.Get().Log.ForceColors,
	})

	Logger.SetOutput(io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   config.Get().Log.File,
		MaxSize:    config.Get().Log.MaxAge,
		MaxBackups: config.Get().Log.MaxBackup,
		MaxAge:     config.Get().Log.MaxAge,
		Compress:   config.Get().Log.Compress,
	}))
}
