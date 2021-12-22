package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	convertType()
	fmt.Println("finish")
}

func convertType() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic：", r)
			debug.PrintStack() // 输出调用站
		}
	}()
	var a interface{} // 1.18 any 可能是任意一种类型
	a = "string aaa"

	b := a.(int)
	fmt.Println(b) //panic: interface conversion: interface {} is string, not int
}
