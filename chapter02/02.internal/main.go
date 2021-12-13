package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	// arr := []int{}
	// 运行返回的时候执行的代码
	startTime := time.Now()
	defer func() { // defer 可以有很多个，从后往前执行
		finishTime := time.Now()
		fmt.Println(startTime, finishTime)
		fmt.Println(finishTime.Sub(startTime))
	}()
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
	}()
	arr2 := testPanic()

	arr3 := make([]string, 3, 4)
	copy(arr3, arr2)
	fmt.Println("复制的arr2：", arr3)

	i := new(int)                  //func new(Type) *Type
	fmt.Println(reflect.TypeOf(i)) //*int

	// 预告
	j := 3333
	k := &j
	fmt.Println(reflect.TypeOf(k))
}

func testPanic() []string {
	arr2 := make([]string, 0, 4)
	fmt.Println("len :", len(arr2), ",cap:", cap(arr2))
	fmt.Println("default:", arr2[0])
	fmt.Println("default:", arr2[1])
	fmt.Println("default:", arr2[2])
	return arr2
}
