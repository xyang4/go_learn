package main

import "fmt"

/**
 函数练习
	基本语法：
		func function_name( [parameter list] ) [return_types] {
		   函数体
		}
*/

func main() {
	baseFunc()
}
func baseFunc() {
	var a, b = 100, 200
	r := max(a, b)
	fmt.Printf("max[%d]", r)
}

// 返回最大值
func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
