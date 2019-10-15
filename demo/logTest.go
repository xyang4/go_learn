package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

/**
  标准库 - log
*/
var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	//log.SetPrefix("TRACE:")
	//log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	file, err := os.OpenFile("error.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Trace = log.New(ioutil.Discard,
		"TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout,
		"INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout,
		"WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr),
		"ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

}

func main() {
	//log.Println("msg")
	//log.Fatalln("fatal msg")
	//log.Panicf("panicf msg")
	Trace.Println("1")
	Info.Println("2")
	Warning.Println("3")
	Error.Println("4")
}
