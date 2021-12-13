package main

import (
	"fmt"
	"reflect"
)

/**
reflect.Typeof 获取类型
*/
func main() {
	a1 := "hello"
	fmt.Println(reflect.TypeOf(a1)) // string
	a2 := 3
	fmt.Println(reflect.TypeOf(a2)) // int
	a3 := 3.0
	fmt.Println(reflect.TypeOf(a3)) // float64
	a4 := &a3
	fmt.Println(a4)                 // 0xc0000a2080 内存地址
	fmt.Println(reflect.TypeOf(a4)) // *float64
}
