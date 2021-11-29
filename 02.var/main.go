package main

import (
	"fmt"
	"math"
)

var i1 = 9

func main() {
	var hello string = "hello goland!" // 编辑器认为是string类型，string可以忽略
	var world = "world"                // 猜测数据类型
	fmt.Println(hello, world)
	小数 := 1.234 // float64
	fmt.Println(小数)

	// 多种方式定义多个变量
	int1, int2 := 4.0, "123"
	fmt.Println(int1, int2)
	var int3, int4 uint = 333, 444 // 编辑器认为是int64 ，与想要的类型不符合！
	fmt.Println(int3, int4)
	var (
		int5, int6 = 333, 222
	)
	fmt.Println(int5, int6)

	var a int
	var b float64
	var isTest bool
	fmt.Println(a)      // 0
	fmt.Println(b)      // 0
	fmt.Println(isTest) // false
	var newname string
	fmt.Println("newname = ", newname)    //newname =
	fmt.Printf("newname = %q\n", newname) //newname = ""

	// 数字之间可以强制转换，但可能会损失精度
	var f1 float64 = 1.234
	fmt.Printf("%p\n", &i1) //0x817250
	fmt.Println("main方法外部的i1：", i1)
	var i1 int = int(f1)
	fmt.Printf("%p\n ", &i1)           //0xc0000a20b0
	fmt.Println("f1：", f1, " i1:", i1) // f1： 1.234  i1: 1
	var int7 uint = math.MaxUint64
	fmt.Println(int7)        //18446744073709551615
	var int8 int = int(int7) // 无符号转有符号
	fmt.Println(int8)        // -1

	// 变量不能重名   ,一个作用域内不能重名
	//i1 := 1 //No new variables on the left side of ':='
	i1 = 2                   // 变量能重写
	fmt.Printf("%p\n ", &i1) //0xc0000a20b0
	fmt.Println()
	{
		//重写时，新作用域内的变量类型如与外部一致，则地址为原来的地址，类型不一致，创建新的地址
		//i1 = 1
		//fmt.Printf("%p\n ",&i1) //0xc0000a20d0
		i1 := 1.10
		fmt.Printf("%p\n ", &i1) //0xc00000a128
		fmt.Println("新作用域中的i1：", i1)
	}

	i1, i2 := 12, "你好"
	fmt.Println(i1, i2)
	fmt.Printf("%p\n ", &i1) //0xc0000a20b0
}
