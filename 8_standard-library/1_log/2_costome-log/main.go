package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger // Just about anything, bcs will be ignored
	Info    *log.Logger // Important Information
	Warning *log.Logger // Be concerned
	Error   *log.Logger // Critical Problem
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file: ", err)
	}

	logFlag := log.Ldate | log.Lmicroseconds | log.Llongfile
	Trace = log.New(ioutil.Discard, "TRACE: ", logFlag)
	Info = log.New(os.Stdout, "INFO: ", logFlag)
	Warning = log.New(os.Stdout, "WARN: ", logFlag)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", logFlag)
}

func main() {
	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}
