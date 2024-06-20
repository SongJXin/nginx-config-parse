package util

import "github.com/sirupsen/logrus"

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	// 设置TextFormatter并启用颜色
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
	Logger.SetLevel(logrus.DebugLevel)
}
