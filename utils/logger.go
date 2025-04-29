package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()
	Log.SetFormatter(&logrus.JSONFormatter{}) // or TextFormatter
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.InfoLevel)
}
