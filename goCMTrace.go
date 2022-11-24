package goCMTrace

import (
	"log"
	"os"
	"time"
)

type logEntry struct {
	Message   string
	time      time.Time
	date      time.Time
	Component string
	Context   string
	State     int32
	Thread    string
	Tile      string
}

func logData(logLine logEntry) error {
	logFile, err := os.OpenFile(logLine.file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	info := "<![LOG[" + logLine.message + "]LOG]!><time=\"" + logLine.time.Local().String() + "\" date=\"" + logLine.date.Local().String() + "\" component=\"CoreSettings\" context=\"\" type=\"1\" thread=\"1\" file=" + logLine.file + "\"\">\n"
	defer logFile.Close()
	if _, err := logFile.WriteString(info); err != nil {
		log.Println(err)
	}
	return err
}
