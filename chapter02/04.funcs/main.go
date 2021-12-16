package main

import (
	"fmt"
	"time"
)

var tall float64
var weight float64

func main() {
	panicAndRecover()

	deferGuess()

	//open /test.txt: The system cannot find the file specified.
	openfile()
	closuerMain()

	// close call 闭包
	fmt.Println(calcSum(1, 123, 43, 21, 5, 52))
	fmt.Println(calcSum(1, 123, 43, 21, 5, 52))
	fmt.Println(calcSum(1, 123, 43, 21, 5, 52))
	showUseTimes()

	guess(1, 100)

	fmt.Println(fib(100)) //随着n的值变大，资源真用越来越高  复杂度是2的n次方
	time.Sleep(10 * time.Second)

	fmt.Println("全局变量赋值前：")
	calcAdd() // 0
	tall, weight = 1.70, 70.0
	fmt.Println("全局变量赋值后：")
	calcAdd() // 71.7
	tall, weight := 100, 70.0
	calcAdd() // 71.7
	tall, weight = 200, 70.0
	calcAdd() // 71.7

	simpleSubdomain()
	// 作用域不同
	//tall, weight := 70, 70.0

	calculateAdd := func(a, b float64) float64 {
		return a + b
	}
	result := calculateAdd(1, 2)
	fmt.Println(result)

	{
		//fmt.Scanln
		person1Tall := 1.81
		person1Weight := 90.0
		calculateAdd(person1Tall, person1Weight)
		//suggestions
	}

	{
		//fmt.Scanln
		person1Tall := 1.81
		person1Weight := 90.0
		calculateAdd(person1Tall, person1Weight)
		//suggestions
	}

	fmt.Println(tall, weight)
}

func calcAdd() int {
	// 作用域不同
	//tall , weight := "你好",70.0
	fmt.Println(tall + weight)
	return 0
}

func simpleSubdomain() {
	name := "小强"
	fmt.Println("名字是：", name)
	{
		fmt.Println("名字是：", name)
		//name = "Kr"
		//fmt.Println("名字是：" ,name) //Kr
		name := "Kr"
		fmt.Println("名字是：", name) //Kr
	}
	fmt.Println("名字是：", name) // 小强
}

func sampleSubdomainIf() {
	if v := calcAdd(); v == 6 {

	} else {

	}
	//fmt.Println(v)//v无效,v的有效范围是if block

	//for中定义的变量也是一样，有效范围都是for block
}

func fib(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func funDef(nums ...int) (addResult int) {
	for _, i := range nums {
		addResult += i
	}
	return //默认返回addResult变量
}
