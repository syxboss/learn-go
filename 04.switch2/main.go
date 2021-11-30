package main

import (
	"fmt"
	"reflect"
)

func main() {
	var money interface{} = 10
	//var money interface{} = 10.0
	//var money interface{} = "10"
	switch newmoney := money.(type) {
	case int:
		var tmpmoney = newmoney + 10.00 // ??
		fmt.Println("money是int类型", newmoney+3.000, tmpmoney)
	case int64:
		fmt.Println("money是int64类型")
	case float64:
		fmt.Println("money是float64类型")
	case float32:
		fmt.Println("money是float32类型")
	default:
		fmt.Println("money是未知类型")
	}
	fmt.Println("结束", money)

	fmt.Println(reflect.TypeOf(money))
	money = 2
	fmt.Println(reflect.TypeOf(money))
}
