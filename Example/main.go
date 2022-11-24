package main

import (
	"github.com/JordanTheITGuy/goCMTrace"
)

func main() {

	logMessage := "Hello World"
	logObj := new(goCMTrace.LogEntry)
	logObj.Message = logMessage
	logObj.Context = "hello World"
	logObj.State = 2
	logObj.File = "helloworld.log"
	goCMTrace.LogData(*logObj)

}
