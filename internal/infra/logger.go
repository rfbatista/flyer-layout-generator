package infra

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewFileLogger() {}

func NewLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	// fullpath := filepath.Join(FindProjectRoot(), "proxy.log")
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// config.Level = zap.NewAtomicLevel()
	// config.OutputPaths = []string{
	// 	fullpath,
	// }
	logger, _ := config.Build()
	return logger
}

func NewTestLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()
	return logger
}
