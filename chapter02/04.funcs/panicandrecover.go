package main

import "fmt"

func panicAndRecover() {
	var nameScore map[string]int = nil
	//defer coverPanicUpgrade()
	defer func() { // 不能捕获异常
		coverPanicUpgrade()
	}()
	//panic: assignment to entry in nil map
	nameScore["姜文"] = 100
}

//和之前defer一样，最终结果是报错
func coverPanic() {
	func() {
		if r := recover(); r != nil {
			fmt.Println("系统出现严重故障")
		}
	}()
}

//捕获异常
func coverPanicUpgrade() {
	if r := recover(); r != nil {
		fmt.Println("系统出现严重故障")
	}
}
