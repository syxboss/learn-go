package main

import "fmt"

func guess(left, right uint) {
	guessed := (right + left) / 2
	var getFromImput string
	fmt.Println("我猜的是：", guessed)
	fmt.Println("如果高了，输入1，如果低了，输入0；对了，输入9")
	fmt.Scan(&getFromImput)
	switch getFromImput {
	case "1":
		if left == right {
			fmt.Println("你是不是改变主意了")
			return
		}
		guess(left, right-1)
	case "0":
		guess(left+1, right)
	case "9":
		fmt.Println("你心里的数字是：", guessed)
	}
}
