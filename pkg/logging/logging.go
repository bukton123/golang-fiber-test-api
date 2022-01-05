package logging

import (
	"go.uber.org/zap"
	"log"
	"os"
)

var logger *zap.Logger

func New(service string, version string) {
	var err error

	host, err := os.Hostname()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	cfg := zap.NewProductionConfig()
	cfg.InitialFields = map[string]interface{}{
		"service":  service,
		"version":  version,
		"hostname": host,
	}

	logger, err = cfg.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
}

func Close() {
	if logger != nil {
		logger.Sync()
	}
}

func Info(message string) {
	logger.Error(message)
}

func Error(message string) {
	logger.Error(message)
}

func Warn(message string) {
	logger.Warn(message)
}

func Fatal(message string) {
	logger.Fatal(message)
}

func output(service string, version string) []zap.Field {
	return []zap.Field{
		zap.String("service", service),
		zap.String("version", version),
	}
}
