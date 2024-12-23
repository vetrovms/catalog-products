package helpers

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	Log.SetFormatter(&logrus.JSONFormatter{})
	if !testing.Testing() {
		path, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		file, err := os.OpenFile(path+"/log/log.json", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			panic(err)
		}
		Log.SetOutput(file)
	}
}
