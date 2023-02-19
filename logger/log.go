package logger

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

// use zerolog for logging. file name will be date in YYYY-MM-DD.json format
func InitLogger() *zerolog.Logger {
	file, err := os.OpenFile("logs"+time.Now().Format("2006-01-02")+".json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	zerolog.TimeFieldFormat = time.RFC3339
	logger := zerolog.New(file).With().Timestamp().Logger()
	return &logger
}
