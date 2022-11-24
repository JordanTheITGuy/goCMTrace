package main

import (
	"github.com/JordanTheITGuy/goCMTrace"
)

func main() {

	logMessage := "Hello World"
	logObj := new(goCMTrace.LogEntry)
	logObj.Message = logMessage
	logObj.Component = "Example"
	logObj.Context = "hello World"
	logObj.State = 2
	logObj.Thread = "hi"
	logObj.File = "helloworld.log"
	goCMTrace.LogData(*logObj)

}
