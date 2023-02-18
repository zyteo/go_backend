package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var Logger *zap.Logger

func InitLogger() *zap.Logger {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncode := zapcore.NewJSONEncoder(config)
	file, _ := os.OpenFile("logs"+time.Now().Format("2006-01-02")+".json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(file)
	defaultLogLevel := zapcore.InfoLevel
	coreInstance := zapcore.NewCore(fileEncode, writer, defaultLogLevel)
	logger := zap.New(coreInstance, zap.AddCaller())
	return logger
}
