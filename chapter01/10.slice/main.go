package main

import "fmt"

func main() {
	// 不能做到任意的转换，只支持有限的转换
	a := "hello" // 默认为UTF8
	a = "你好"
	fmt.Println(a)
	aBytes := []byte(a)
	fmt.Println(aBytes)
	fmt.Println("修改切片内的内容")
	aBytes[0] = 'H'
	a = string(aBytes)
	fmt.Println(a) // 长度为6

	b := "你好"
	fmt.Println(b)
	aBytes1 := []rune(b) //字符串每个字符当成是是单个字符
	fmt.Println(aBytes1)
	aBytes1[0] = 'H'
	b = string(aBytes1)
	fmt.Println(b) //H好  长度为2
}
