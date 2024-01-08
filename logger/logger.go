package logger

import "go.uber.org/zap"

var Logger *zap.Logger

func Init() error {
	var err error
	Logger, err = zap.NewDevelopment()
	if err != nil {
		return err
	}
	return nil
}

func GetLogger() *zap.Logger {
	return Logger
}
