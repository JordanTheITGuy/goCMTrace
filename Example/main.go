package main

import (
	"github.com/JordanTheITGuy/goCMTrace"
)

func main() {

	logMessage := "Hello World state 3"
	logObj := new(goCMTrace.LogEntry)
	logObj.Message = logMessage
	logObj.State = 3
	logObj.File = "helloworld.log"
	goCMTrace.LogData(*logObj)

}
