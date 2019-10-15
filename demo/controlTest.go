package main

import "fmt"

func main() {
	// 一旦成功匹配到 case（无 fallthrough 说明）便直接退出，无需 break 来结束
	switchFunc()
}
func switchFunc() {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
