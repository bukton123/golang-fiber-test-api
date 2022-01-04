package logging

import (
	"go.uber.org/zap"
	"log"
)

var logger *zap.Logger

func New(service string, version string) *zap.Logger {
	var err error
	cfg := zap.Config{
		Encoding: "json",
		Level:    zap.NewAtomicLevelAt(zap.DebugLevel),
		InitialFields: map[string]interface{}{
			"service": service,
			"version": version,
		}}

	b, err := cfg.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)

	}

	return b
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

func output(service string, version string) []zap.Field {
	return []zap.Field{
		zap.String("service", service),
		zap.String("version", version),
	}
}
