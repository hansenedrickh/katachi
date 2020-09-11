package common

import (
	"fmt"
	"os"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func setupLogger() {
	formatter := runtime.Formatter{ChildFormatter: &logrus.JSONFormatter{}}
	formatter.Line = true
	formatter.File = true

	Log = &logrus.Logger{
		Out:       os.Stderr,
		Formatter: &formatter,
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
}

func BuildContext(context string, fields map[string]interface{}) logrus.Fields {
	logFields := logrus.Fields{
		"context": context,
	}

	for key, value := range fields {
		logFields[key] = value
	}

	return logFields
}

func LogError(fields logrus.Fields, err error) {
	msg := fmt.Sprintf("got error for %s: %s", fields["request-id"], err.Error())
	Keisatsu.Error(msg)
	Log.WithFields(fields).Error(err.Error())
}

func LogInfo(fields logrus.Fields, msg string) {
	Log.WithFields(fields).Error(msg)
}




