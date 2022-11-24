package main

import (
	"github.com/JordanTheITGuy/goCMTrace"
)

func main() {

	logMessage := "Hello World"
	log := new(goCMTrace.LogEntry)
	log.Message = logMessage
	log.Component = "Example"
	log.Context = "hello World"
	log.State = 2
	log.Thread = "hi"
	log.File = "helloworld.log"

}
