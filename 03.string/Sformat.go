package main

import "fmt"

func main() {
	fmt.Println("中文，123 ，abc，❤")

	fmt.Printf("我的名字是%s\n", "小轩轩") //我的名字是小轩轩
	fmt.Printf("我的名字是%q\n", "小轩轩") //我的名字是"小轩轩"
	fmt.Printf("我的名字是%x\n", "小轩轩") //我的名字是e5b08fe8bda9e8bda9  二字符十六进制（a-f）
	fmt.Printf("我的名字是%X\n", "小轩轩") //我的名字是E5B08FE8BDA9E8BDA9	二字符十六进制（A-F）

}
