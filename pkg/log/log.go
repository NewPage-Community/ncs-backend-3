package log

import (
	"fmt"

	"go.uber.org/zap"
)

const callerSkip = 1

type Config struct {
	Debug bool
}

var _defConfig = &Config{
	Debug: false,
}

var logger *zap.Logger

func Init(conf *Config) {
	if conf == nil {
		conf = _defConfig
	}

	var err error
	if conf.Debug {
		logger, err = zap.NewDevelopment(zap.AddCaller(), zap.AddCallerSkip(callerSkip))
	} else {
		logger, err = zap.NewProduction(zap.AddCaller(), zap.AddCallerSkip(callerSkip))
	}

	if err != nil {
		panic(err)
	}
}

func Close() {
	err := logger.Sync()
	if err != nil {
		fmt.Println("zap error:", err.Error())
	}
}

func Info(args ...interface{}) {
	logger.Sugar().Info(args)
}

func Error(args ...interface{}) {
	logger.Sugar().Error(args)
}

func Warn(args ...interface{}) {
	logger.Sugar().Warn(args)
}

func Debug(args ...interface{}) {
	logger.Sugar().Debug(args)
}

func Panic(args ...interface{}) {
	logger.Sugar().Panic(args)
}

func CheckErr(err error) {
	if err != nil {
		logger.Sugar().Info(err)
	}
}

func GetLogger() *zap.Logger {
	return logger
}
