package main

import (
	"fmt"
)

func main() {
	var hello string = "Hello , golang !"
	fmt.Println(hello)
	var age int = 35
	fmt.Println(age)
	var tall float64 = 1.80
	fmt.Println(tall)

	//报错信息：'12' (type untyped int) cannot be represented by the type string
	//hello = 12 // 类型不兼容
	//Cannot use 'age' (type int) as the type float64
	//tall = age //类型不兼容

}
