package main

// 结构体嵌套
type Newcalculator struct {
	Calculator // 对象会自动实例化，指针类型则不会
	//*Calculator // 继承指针很危险,未实例化就会报错
}
