package main

import "fmt"

func main() {
	var money int = 20
	var busy bool = false
	switch money {
	case 20:
		fmt.Println("点个外卖")
		if busy {
			break
		} else {
			fmt.Println("再吃个零食")
		}
		fallthrough // 直接并入下一个处理分支而无需判断条件
	case 200:
		fmt.Println("下个馆子")
	default:
		fmt.Println("容我想想")
	}
	fmt.Println("end")
}
