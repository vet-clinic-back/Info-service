package logging

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Entry
}

func NewLogger(isLocal, isDebug *bool) *Logger {
	resLogger := &Logger{}

	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{}
	log.SetLevel(logrus.InfoLevel)

	resLogger.Entry = logrus.NewEntry(log)

	if *isLocal {
		log.Formatter = &logrus.TextFormatter{
			ForceColors:      true,
			DisableTimestamp: false,
			FullTimestamp:    true,
		}
		resLogger.Entry = logrus.NewEntry(log)
	}

	if *isDebug {
		resLogger.Logger.SetLevel(logrus.DebugLevel)
	}

	return resLogger
}
