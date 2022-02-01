package main

import (
	"fmt"
	"testing"
)

func TestAssertion(t *testing.T) {
	r := TestBox{}
	var c Close = r

	// 第一种断言
	switch cDetail := c.(type) {
	case Refrigerator:
		fmt.Println("是预期的冰箱")
		fmt.Println(cDetail.Size)
	case Box:
		fmt.Println("这是一个box，不能当成冰箱使")
	}
	//r2 := c.(Refrigerator)  //main.Close is *main.TestBox, not main.Refrigerator
	//fmt.Println(r2.Size)

	{
		refrigerator, ok := checkIsRefrigerator(c)
		if ok {
			fmt.Println("是个冰箱，开门装大象！", refrigerator)
		} else {
			fmt.Println("不是个冰箱")
		}
	}
}

// 第二种断言
//如果断言成功，Value的值为i的值;如果失败，则为默认值
//如果断言成功，ok的值为true;否则为false
//i为需要断言的类型
//.(<type>)为断言格式。<type>为尝试断言的类型

func checkIsRefrigerator(c Close) (Refrigerator, bool) {
	r, ok := c.(Refrigerator)
	return r, ok
}

type TestBox struct {
}

func (tb TestBox) Close() error {
	return nil
}
