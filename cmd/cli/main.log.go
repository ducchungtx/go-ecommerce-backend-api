package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// 1.
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name:%s, age:%d", "John", 30)

	// 2.

	// logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "John"), zap.Int("age", 30))

	// logger := zap.NewExample()
	// logger.Info("Hello World")

	// Development logger
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello World Development")

	// Production logger
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello World Production")

	// 3.
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // ISO8601TimeEncoder: 2021-08-26T15:04:05.999+07:00
	encoderConfig.TimeKey = "time"                          // time
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // CapitalLevelEncoder: INFO, ERROR, DEBUG
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder // ShortCallerEncoder: /path/to/file.go:line
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	// Ensure the directory exists
	if _, err := os.Stat("./log"); os.IsNotExist(err) {
		err := os.Mkdir("./log", os.ModePerm)
		if err != nil {
			panic("Failed to create log directory: " + err.Error())
		}
	}

	file, err := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}

	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
