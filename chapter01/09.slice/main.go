package main

import "fmt"

func main() {
	a := []int{} // 切片
	fmt.Println(a)
	fmt.Println("追加元素a中，a是切片")
	a = append(a, 3333)
	fmt.Println(a) // 动态增加
	fmt.Println("删除切片a中的元素")
	a = []int{111, 222, 333, 444, 555}
	fmt.Println("删除前：", a)
	a = append(a[:2], a[3:]...)
	fmt.Println("删除后：", a)
	a = append(a, a[:]...)
	fmt.Println("double:", a)

	backup := a[1:] // 同一个数组的指针  111 999 999 444 555 111 222 444 555
	backup = append([]int{}, a[1:]...)
	a = append(a[:1], 999)
	a = append(a, backup...)
	fmt.Println("添加：", a) //111 999 444 555 111 222 444 555

	b := [0]int{} // 数组  没有增加删除的方法
	fmt.Println("追加元素b中，b是数组")
	//b = append(b,3333) //Cannot use 'b' (type [0]int) as the type []Type
	fmt.Println(b)
}
