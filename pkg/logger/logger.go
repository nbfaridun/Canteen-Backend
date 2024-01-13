package logger

import "go.uber.org/zap"

var Logger *zap.Logger

func Init() error {
	var err error
	Logger, err = zap.NewDevelopment(zap.AddCallerSkip(1), zap.AddStacktrace(zap.DPanicLevel))
	if err != nil {
		return err
	}
	return nil
}

func GetLogger() *zap.Logger {
	return Logger
}
