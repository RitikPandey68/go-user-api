package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

// NewLogger creates and returns a new production-grade Uber Zap logger.
// It uses JSON encoding in production for structured log output.
func NewLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	var err error
	log, err = config.Build()
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
	return log
}

// Info logs a message at INFO level with optional structured fields.
func Info(msg string, fields ...zap.Field) {
	if log != nil {
		log.Info(msg, fields...)
	}
}

// Error logs a message at ERROR level with optional structured fields.
func Error(msg string, fields ...zap.Field) {
	if log != nil {
		log.Error(msg, fields...)
	}
}

// Warn logs a message at WARN level with optional structured fields.
func Warn(msg string, fields ...zap.Field) {
	if log != nil {
		log.Warn(msg, fields...)
	}
}

// Sync flushes any buffered log entries. Call before application exit.
func Sync() {
	if log != nil {
		_ = log.Sync()
	}
}
