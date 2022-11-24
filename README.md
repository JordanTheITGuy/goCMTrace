# goCMTrace

A Go Module for outputing results in the popular CMTrace Format. 


## Usage

If you've ever worked with Configuration Manager Endpoint Manager at some point you've worked with CMTrace, and you probably really wish that other languages would allow you to use this log method. Well now you can use it with Go. This package is publicly available and could be used in any small GO project that is designed to run locally on a machine. I would not suggest using this for larger projects or API's. 

By Default, all log messages will default to a stadard message if you want to create a warning, or an error type be sure to update the STATE of the object. See Examples for details.

## How to import this package

This module is available on GitHub, and you can import it directly from GitHub by including it in your import statement.

```go
import(
    "github.com/JordanTheITGuy/goCMTrace"
)
```
## Example

There is a working example of a main.go file in the Example folder, and a screenshot of the output see below for comments and details. 

![CMTrace Example](/Example/exampleOutPut.png)

```go
    //Set the message you would like to log.
	
    logMessage := "Hello World"
    //Create the Log Entry Object.

	logObj := new(goCMTrace.LogEntry)

    //Set the message state on the newly created log object. You can just recycle the log object, you do not need multiple log objects.

	logObj.Message = logMessage

    //Specify the log file you would like to update, you can change it multiple times if you want.

	logObj.File = "helloworld.log"
    //Send the lob object through to the LogData Function.

    //By default, if you do not set or provide a state, it is assumed it should log in state 1. Make sure you revert back, to 1 if you send a log message through. 

    logObj.State = "2"

	goCMTrace.LogData(*logObj)
```

