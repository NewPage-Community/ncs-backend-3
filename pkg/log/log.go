package log

import "go.uber.org/zap"

var logger *zap.Logger

func Init(debug bool) {
	var err error
	if debug {
		logger, err = zap.NewDevelopment(zap.AddCallerSkip(1))
	} else {
		logger, err = zap.NewProduction(zap.AddCallerSkip(1))
	}

	if err != nil {
		panic("Panic: " + err.Error())
	}
}

func Close() {
	logger.Sync()
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