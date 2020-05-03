package log

import (
	"fmt"
	"go.uber.org/zap"
)

type Config struct {
	Debug bool
}

var logger *zap.Logger

func Init(conf *Config) {
	var err error
	if conf.Debug {
		logger, err = zap.NewDevelopment(zap.AddCallerSkip(1))
	} else {
		logger, err = zap.NewProduction(zap.AddCallerSkip(1))
	}

	if err != nil {
		panic("Panic: " + err.Error())
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
