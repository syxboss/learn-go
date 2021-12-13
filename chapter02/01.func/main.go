package main

import "fmt"

func main() {
	hello()
	helloToSomeOne("Dandy")
	msg := constructHelloMessage("张三")
	fmt.Println(msg)
}

func hello() {
	fmt.Println("你好Golang！")
}

func helloToSomeOne(name string) {
	fmt.Println("你好，", name)
}

func constructHelloMessage(name string) string {
	return "你好，" + name
}
