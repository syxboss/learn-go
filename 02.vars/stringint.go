package main

import (
	"fmt"
	"strconv"
)

func main() {
	// string和int之间的转换
	a1 := 100
	s1 := "122"
	s2 := "你好"
	// strconv
	// 1.Atoi
	var value, err = strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", value, value) //type:int value:122
	}

	var value2, err2 = strconv.Atoi(s2)
	if err2 != nil {
		fmt.Println(value2, err2) //0  strconv.Atoi: parsing "你好": invalid syntax
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", value2, value2) //type:int value:100
	}

	// 2.itoa
	var value3 = strconv.Itoa(a1)                     //返回 string类型
	fmt.Printf("type:%T value:%#v\n", value3, value3) //type:string value:"100"

}
