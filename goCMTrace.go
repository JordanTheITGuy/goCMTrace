package goCMTrace

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

type LogEntry struct {
	Message   string
	time      time.Time
	date      time.Time
	Component string
	Context   string
	State     int
	Thread    string
	File      string
}

func LogData(logLine LogEntry) error {
	logFile, err := os.OpenFile(logLine.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		logLine.Component = details.Name()
		callingFile, callingLine := details.FileLine(pc)
		logLine.Thread = callingFile + " " + strconv.Itoa(callingLine)
	}
	info := "<![LOG[" + logLine.Message + "]LOG]!><time=\"" + logLine.time.Local().String() + "\" date=\"" + logLine.date.Local().String() + "\" component=\"" + logLine.Component + "\" context=\"" + logLine.Context + "\" type=\"" + strconv.Itoa(logLine.State) + "\" thread=\"" + logLine.Thread + "\" file=" + logLine.File + "\"\">\n"
	defer logFile.Close()
	if _, err := logFile.WriteString(info); err != nil {
		log.Println(err)
	}
	return err
}
