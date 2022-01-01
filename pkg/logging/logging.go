package logging

import (
	"go.uber.org/zap"
	"log"
)

var logger *zap.Logger

func init() {
	if logger == nil {
		var err error
		logger, err = zap.NewProduction()
		if err != nil {
			log.Fatalf("can't initialize zap logger: %v", err)
		}

		defer logger.Sync()
	}
}

func Info(service string, version string, message string) {
	logger.Error(message, output(service, version)...)
}

func Error(service string, version string, message string) {
	logger.Error(message, output(service, version)...)
}

func Warn(service string, version string, message string) {
	logger.Warn(message, output(service, version)...)
}

func output(service string, version string) []zap.Field {
	return []zap.Field{
		zap.String("service", service),
		zap.String("version", version),
	}
}
