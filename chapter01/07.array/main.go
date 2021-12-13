package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	var b [3]int = [3]int{}
	//var c [3]int = [4]int{} //Cannot use '[4]int{}' (type [4]int) as the type [3]int
	d := [5]int{}
	e := [...]int{1, 2, 3, 4} // 数组...是什么意思？ 长度不需要管理，有多少就是多少
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println("======================")

	// 动态赋值
	// var xxx type 固定类型有初值
	var k [3]int
	//k = [3]string{"","",""} //Cannot use '[3]string{"","",""}' (type [3]string) as the type [3]int
	k = [3]int{1, 1, 1}
	fmt.Println(k)
	k = [3]int{2, 2, 2}
	fmt.Println(k)
	fmt.Println("======================")

	// 数组赋值
	z := [3]int{}
	fmt.Println(z)
	z[0] = 1
	z[1] = 2
	z[2] = 3
	//z[4] = 4 // 编译不通过Invalid array index '4' (out of bounds for the 3-element array）越界
	fmt.Println(z)
	fmt.Println("======================")

	//数组长度 len获取长度
	l := [3]int{}
	fmt.Println("数组l的长度为", len(l))
	fmt.Println("======================")

	// 遍历数组元素
	o := [4]int{1, 2, 3, 4}
	for i := 0; i < len(o); i++ {
		fmt.Println("使用下标进行打印：", o[i])
	}

	for index, item := range o {
		fmt.Printf("打印数组下标为：%d,值为：%d \n", index, item)
	}
	fmt.Println("======================")

	//var num *int // 指针类型，没有指向变量的时候就会直接退出

	// 多维数组
	p := [3][3]int{}
	p[0] = [3]int{1, 2, 3}
	p[1] = [3]int{4, 5, 6}
	p[2] = [3]int{7, 8, 9}

	p1 := [...][3]int{
		[3]int{1, 2, 3},
		[3]int{4, 5, 6},
		[3]int{7, 8, 9},
	}
	fmt.Println("p1原样打印：", p1)
	fmt.Println("p原样打印：", p)

	// 多维数组遍历操作
	for i := 0; i < len(p); i++ {
		fmt.Println("打印第一层:", p[i])
	}

	// 多层遍历，每次遍历降低一次维度
	for i, val := range p {
		for index, item := range val {
			fmt.Printf("这是第一层第%d个，第二层第%d个，值为%d\n", i+1, index+1, item)
		}
	}

	// 数组不支持运算，比如数组之间的加法
	aa := []int{} // 切片
	fmt.Println(aa)
	bb := [0]int{} // 数组
	fmt.Println(bb)
}
