package main

import "fmt"

/**
函数练习：
	6.2.1 按值传递（call by value）按引用传递（call by reference）
		- 1 指针传递允许直接修改变量的值，消耗也更少
	6.2.2 命名的返回值（named return variables）
	6.2.3 空白符（blank identifier）
	6.2.4 改变外部变量（outside variable）
*/

func main() {
	multipleReturnFunc()
}

func multipleReturnFunc() {
	num := 10
	var numx2, numx3 int
	numx2, numx3 = getX2AndX3(num)
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
	numx2, numx3 = getX2AndX3_2(num)
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

func getX2AndX3(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 2 * input
	return
}
func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 2 * input
	return
}
