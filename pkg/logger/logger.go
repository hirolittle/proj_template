package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"proj_template/pkg/config"
	"strings"
)

// logger and SugarLogger are global instances for logging
var (
	logger      *zap.Logger
	SugarLogger *zap.SugaredLogger
)

// InitLogger initializes the logger with the specified configuration
func InitLogger() {
	logLevel := getLogLevel(config.Cfg.LogConfig.Level)
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)

	logger = zap.New(core, zap.AddCaller())
	SugarLogger = logger.Sugar()
}

// getLogLevel converts a string log level to zapcore.Level
func getLogLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		log.Printf("Unknown log level: %s, defaulting to Info level", level)
		return zapcore.InfoLevel
	}
}

// getLogWriter sets up the log writer using lumberjack for log rotation
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", config.Cfg.LogConfig.Output, config.Cfg.LogConfig.FileName),
		MaxSize:    config.Cfg.LogConfig.MaxSize,
		MaxBackups: config.Cfg.LogConfig.MaxBackups,
		MaxAge:     config.Cfg.LogConfig.MaxAge,
		Compress:   config.Cfg.LogConfig.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// getEncoder sets up the log encoder with the desired format
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
