package main

import "fmt"

func main() {
	{
		a := make([]int, 0) // len,cap
		fmt.Println("len :", len(a))
		a = append(a, 1)
		fmt.Println("len :", len(a), "val:", a)
	}
	{
		a := make([]int, 1) // len,cap
		fmt.Println("len :", len(a))
		a = append(a, 1)
		fmt.Println("len :", len(a), "val:", a)
	}

	{
		a := make([]int, 0, 1) // len,cap
		fmt.Println("len :", len(a), "cap :", cap(a))
		a = append(a, 1)
		fmt.Println("len :", len(a), "val:", a)
	}

	// cap容量，空间不足内部就会扩容，有性能损失，建议有cap预期
}
