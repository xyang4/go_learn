package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	//strRelatedTest()
	// ==== 时间相关
	timeFunc()
}

func strRelatedTest() {
	//4.7.1 前缀和后缀
	//4.7.2 字符串包含关系
	//4.7.3 判断子字符串或字符在父字符串中出现的位置（索引）
	//4.7.4 字符串替换
	//4.7.5 统计字符串出现次数
	countFunc()
	//4.7.6 重复字符串
	repeatFunc()
	//4.7.7 修改字符串大小写
	//4.7.8 修剪字符串
	//4.7.9 分割字符串
	splitFunc()
	//4.7.10 拼接 slice 到字符串
	//4.7.11 从字符串中读取内容
	strReaderFunc()
	//4.7.12 字符串与其它类型的转换
	strConvertFunc()
}
func timeFunc() {
	var week time.Duration
	t := time.Now()
	fmt.Println(t) // e.g. Wed Dec 21 09:52:14 +0100 RST 2011
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	// 21.12.2011
	t = time.Now().UTC()
	fmt.Println(t)          // Wed Dec 21 08:52:14 +0000 UTC 2011
	fmt.Println(time.Now()) // Wed Dec 21 09:52:14 +0100 RST 2011
	// calculating times:
	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	week_from_now := t.Add(week)
	fmt.Println(week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011
	// formatting times:
	fmt.Println(t.Format(time.RFC822))         // 21 Dec 11 0852 UTC
	fmt.Println(t.Format(time.ANSIC))          // Wed Dec 21 08:56:34 2011
	fmt.Println(t.Format("02 Jan 2006 15:04")) // 21 Dec 2011 08:52
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
}
func strReaderFunc() {
	str := "xYang"
	reader := strings.NewReader(str)
	fmt.Printf("len[%d] size[%d]\n", reader.Len(), reader.Size())
}

func splitFunc() {

}

// strings.Count(s, str string) int
func countFunc() {
	var str = "Hello, how is it going, Hugo?"
	var manyG = "gggggggggg"

	fmt.Printf("Number of H's in %s is: ", str)
	fmt.Printf("%d\n", strings.Count(str, "H"))

	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))
}

func repeatFunc() {
	var origS = "Hi there! "
	var newS string

	newS = strings.Repeat(origS, 3)
	fmt.Printf("The new repeated string is: %s\n", newS)
}

func strConvertFunc() {
	var orig string = "666"
	var an int
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
