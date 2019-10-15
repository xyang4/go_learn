package main

import (
	"fmt"
	"strings"
	"time"
)

type Reader1 interface {
	read(rc chan string)
}
type Writer1 interface {
	write(wc chan string)
}

type LogProcess1 struct {
	rc     chan string
	wc     chan string
	reader Reader1
	writer Writer1
}

type ReadFromFile1 struct {
	path string
}
type WriteToInfluxDB1 struct {
	influxDBDsn string
}

// 1 读取模块
func (l *ReadFromFile1) read(rc chan string) {
	line := "message"
	rc <- line
}

// 2 解析模块
func (l *LogProcess1) process() {
	data := <-l.rc
	l.wc <- strings.ToUpper(data)
}

// 3 写入模块
func (l *WriteToInfluxDB1) write(wc chan string) {
	fmt.Printf(<-wc)
}

// simpleMain0.go 优化，抽象读取/写入func,便于扩展
func main() {
	reader := &ReadFromFile1{
		path: "/tep/access.log",
	}

	writer := &WriteToInfluxDB1{
		influxDBDsn: "username=?&password=?",
	}

	lp := &LogProcess1{
		rc:     make(chan string),
		wc:     make(chan string),
		reader: reader,
		writer: writer,
	}

	go lp.reader.read(lp.rc)
	go lp.process()
	go lp.writer.write(lp.wc)

	time.Sleep(2 * time.Second)
}
