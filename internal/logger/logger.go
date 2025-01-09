package logger

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

const logPath = "LOG_PATH"

// Log. Повертає компонент логування
func Log() *logrus.Logger {
	log.SetFormatter(&logrus.JSONFormatter{})
	if !testing.Testing() {
		file, err := os.OpenFile(os.Getenv(logPath), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			panic(err)
		}
		log.SetOutput(file)
	}
	return log
}
