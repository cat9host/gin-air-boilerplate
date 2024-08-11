package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"sync"
)

var Logger *zap.Logger
var mutex = &sync.Mutex{}

func InitializeLogger() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	Logger = logger

	defer func(Logger *zap.Logger) {
		_ = Logger.Sync()
	}(Logger)
}

func out(level zapcore.Level, message string, domain string, fields ...zap.Field) {
	GetLogger().Log(level, fmt.Sprintf("[%s]: %s", domain, message), fields...)
}

// Info logs information
func Info(message string, domain string) {
	out(zapcore.InfoLevel, message, domain)
}

// Warn logs
func Warn(message string, domain string) {
	out(zapcore.WarnLevel, message, domain)
}

// Error logs error
func Error(message string, domain string, fields ...zap.Field) {
	out(zapcore.ErrorLevel, message, domain, fields...)
}

func FatalError(message string, domain string, fields ...zap.Field) {
	out(zapcore.FatalLevel, message, domain, fields...)
}

// Debug should be used only for debug something
func Debug(args ...interface{}) {
	fmt.Println(args...)
}

func GetLogger() *zap.Logger {
	mutex.Lock()
	if Logger == nil {
		InitializeLogger()
	}
	mutex.Unlock()

	return Logger
}
