package main

import (
	"github.com/JordanTheITGuy/goCMTrace"
)

func main() {

	logMessage := "Hello World state"
	logObj := new(goCMTrace.LogEntry)
	logObj.Message = logMessage
	logObj.File = "helloworld.log"
	goCMTrace.LogData(*logObj)

}
