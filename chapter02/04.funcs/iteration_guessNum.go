package main

import "fmt"

func guess(left, right uint) {
	guessed := (right + left) / 2
	var getFromImput string
	fmt.Println("我猜的是：", guessed)
	fmt.Print("如果高了，输入1，如果低了，输入0；对了，输入9：")
	fmt.Scan(&getFromImput)
	switch getFromImput {
	case "1":
		if left == right {
			fmt.Println("你是不是改变主意了")
			return
		}
		guess(left, guessed-1)
	case "0":
		if left == right {
			fmt.Println("你是不是改变主意了")
			return
		}
		guess(guessed+1, right)
	case "9":
		fmt.Println("你心里的数字是：", guessed)
	default:
		fmt.Println("输入有误！请重新输入！")
		guess(left, right)
	}
}
