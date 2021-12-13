package main

import "fmt"

func main() {
	a, b := 100, 31
	fmt.Println(a ^ b) // 异或的交换律
	fmt.Println(b ^ a)

	// 查找一堆数中的唯一不重复的值
	arr := []int{4, 3, 4, 5, 6, 7, 5, 6, 3}
	result := -1 // 初始化 有效值
	for _, item := range arr {
		if result < 0 {
			result = item
		} else {
			result = result ^ item
		}
	}
	fmt.Println(result)

}
