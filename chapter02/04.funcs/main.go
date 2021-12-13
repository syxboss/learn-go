package main

import "fmt"

var tall string
var weight float64

func main() {
	// 作用域不同
	tall, weight := 70, 70.0
	fmt.Println(tall, weight)
}

func calcBMI() {
	// 作用域不同
	//tall , weight := "你好",70.0
	fmt.Println(tall, weight)
}
