package goCMTrace

import (
	"log"
	"os"
	"time"
)

type LogEntry struct {
	Message   string
	time      time.Time
	date      time.Time
	Component string
	Context   string
	State     int32
	Thread    string
	File      string
}

func LogData(logLine LogEntry) error {
	logFile, err := os.OpenFile(logLine.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	info := "<![LOG[" + logLine.Message + "]LOG]!><time=\"" + logLine.time.Local().String() + "\" date=\"" + logLine.date.Local().String() + "\" component=\"" + logLine.Component + "\" context=\"\" type=\"1\" thread=\"1\" file=" + logLine.File + "\"\">\n"
	defer logFile.Close()
	if _, err := logFile.WriteString(info); err != nil {
		log.Println(err)
	}
	return err
}
