package logger

import (
	"go.uber.org/zap"
	"log"
)

func InitializeLogger() (*zap.SugaredLogger, error) {
	zapLogger, errZap := zap.NewProduction()
	if errZap != nil {
		return nil, errZap
	}
	defer func(zapLogger *zap.Logger) {
		err := zapLogger.Sync()
		if err != nil {
			log.Println(err)
		}
	}(zapLogger)
	logger := zapLogger.Sugar()
	return logger, nil
}
