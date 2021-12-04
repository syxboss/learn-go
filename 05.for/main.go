package main

import "fmt"

func main() {
	// i++ = i += 1 = i = i + 1
	for i := 0; i < 5; i++ {
		fmt.Println("你好，Golang！ ")
	}

	fmt.Println("=============================")
	i := 0
	//for true { //死循环
	for i < 10 {
		fmt.Println("你好，Golang！ ")
		i++
	}

	fmt.Println("=============================")
	j := 0
	for {
		fmt.Println("你好，Golang ！")
		j++
		if j > 3 {
			break // 退出循环
		}
	}

	fmt.Println("=============================")
	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			continue // 跳过当前次循环
		}
		fmt.Println("你好，golang!", i)
	}
}
