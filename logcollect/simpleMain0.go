package main

import (
	"fmt"
	"strings"
	"time"
)

type LogProcess struct {
	rc          chan string
	wc          chan string
	path        string
	influxDBDsn string
}

// 1 读取模块
func (l *LogProcess) ReadFromFile() {
	line := "message"
	l.rc <- line
}

// 2 解析模块
func (l *LogProcess) Process() {
	data := <-l.rc
	l.wc <- strings.ToUpper(data)
}

// 3 写入模块
func (l *LogProcess) writeToInfluxDB() {
	fmt.Printf(<-l.wc)
}

// 模拟并发实时处理日志文件
func main() {
	lp := &LogProcess{
		rc:          make(chan string),
		wc:          make(chan string),
		path:        "/tep/access.log",
		influxDBDsn: "username=?&password=?",
	}

	go lp.ReadFromFile()
	go lp.Process()
	go lp.writeToInfluxDB()

	time.Sleep(2 * time.Second)
}
