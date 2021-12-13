package main

import "fmt"

func main() {
	// 类型不一致的时候无法做运算
	var a, b int = 12, 10 //类型选择不准确，在运算过程中会出现问题，导致溢出
	fmt.Println("a + b=", a+b)
	fmt.Println("a - b=", a-b)
	fmt.Println("a * b=", a*b)
	fmt.Println("a / b=", a/b) // 除数不能为0，panic: runtime error: integer divide by zero
	fmt.Println("a % b=", a%b) // 只能int进行运算，float会报错

	fmt.Println(true && false == false)

	fmt.Println(a >= b)
	fmt.Println(a < b)
}
