package main

import "fmt"

/*
 go 语法 for 练习
	1. for init; condition; post { }

	2. for condition { }

	3. for { }
*/

func main() {
	// for 基本语法
	forTest()

	// 迭代str
	strForItemFunc()

	printFunc()

	rangeStrFunc()
}
func strForItemFunc() {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str[ %s ] is: %d\n", str, len(str))
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
	}
	//循环迭代一个 Unicode 编码的字符串
	str2 := "日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}
}

func forTest() {
	var a int
	var b = 15
	numbers := []int{1, 3, 2, 6, 4}

	for a := 0; a < 10; a++ {
		fmt.Printf("a[%d]\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a[%d]\n", a)
	}

	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
}

// 使用 * 符号打印宽为 20，高为 10 的矩形。
func printFunc() {
	w, h := 20, 10
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func rangeStrFunc() {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}
	fmt.Println()
	str2 := "Chinese: 日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, char := range str2 {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
	fmt.Println()
	fmt.Println("index int(rune) rune    char bytes")
	for index, rune := range str2 {
		fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
	}
}
