package logwrapper

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogging() {

	var maxSize = viper.GetInt("logging.maxSizeInMb")
	var maxAge = viper.GetInt("logging.maxAgeInDays")

	var logFile = &lumberjack.Logger{
		Filename: "./logs/log.txt",
		MaxSize:  maxSize,
		MaxAge:   maxAge}

	var mw = io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}
