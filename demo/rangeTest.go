package main

import "fmt"

/**
range 练习
	用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
	在数组和切片中它返回元素的索引和索引对应的值
	在集合中返回 key-value 对的 key 值。
*/
func main() {
	//这是我们使用range去求一个slice的和。
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//在数组上使用range将传入index和值两个变量
	for i, num := range nums {
		fmt.Printf("exe for item index[%d] val[%d]\n", i, num)
	}
	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
