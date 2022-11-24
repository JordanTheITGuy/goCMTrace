package goCMTrace

import (os

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

type logEntry struct {
	message   string
	time      time.Time
	date      time.Time
	component string
	context   string
	state     int32
	thread    string
	file      string
}