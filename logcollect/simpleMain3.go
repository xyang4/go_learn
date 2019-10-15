package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Reader3 interface {
	read(rc chan []byte)
}

type Writer3 interface {
	write(wc chan string)
}
type ReadFromFile3 struct {
	path string
}
type WriteToInfluxDB3 struct {
	influxDBDsn string
}
type LogProcess3 struct {
	rc     chan []byte
	wc     chan string
	reader Reader3
	writer Writer3
}

// 1 读取模块
func (r *ReadFromFile3) read(rc chan []byte) {

	fmt.Println("Begin read File")
	file, e := os.OpenFile(r.path, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if e != nil {
		panic(fmt.Sprintf("open file error:%s", e.Error()))
	}

	// 从文件末尾开始逐行读取文件内容
	file.Seek(0, 2)
	rd := bufio.NewReader(file)

	for {
		line, err := rd.ReadBytes('\n')

		if err != io.EOF { // 到结尾
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("ReadBytes error:%s", err.Error()))
		}
		fmt.Println("> ", line)
		rc <- line
		//rc <- line[:len(line)-1]
	}

}

func (l *LogProcess3) process() {
	// 解析模块
	for v := range l.rc {
		l.wc <- strings.ToUpper(string(v))
	}
}

// 3 写入模块
func (w *WriteToInfluxDB3) write(wc chan string) {
	for v := range wc {
		fmt.Printf(v)
	}
}

func main() {
	read := &ReadFromFile3{
		path: "/Users/xyang/go_code/src/xyang.com/logcollect/data/access.log",
		//path: "data/access.log",
	}

	writer := &WriteToInfluxDB3{
		influxDBDsn: "username=?&password=?",
	}

	lp := &LogProcess3{
		rc:     make(chan []byte),
		wc:     make(chan string),
		reader: read,
		writer: writer,
	}

	go lp.reader.read(lp.rc)
	go lp.process()
	go lp.writer.write(lp.wc)

	time.Sleep(30 * time.Second)
}
