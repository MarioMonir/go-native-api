package logger_utils

import (
	"log"
	"os"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "[api] ", log.LstdFlags)
}
