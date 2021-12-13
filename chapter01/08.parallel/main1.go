package main

import "fmt"

/*
判断两条线是否平行
	提示：
		两点决定一条直线
		两条线是否平行取决于两条线的斜率是否一样
*/
func main() {
	for {
		//斜率
		var k = [2]float64{}
		for i := 0; i < 2; i++ {
			k[i] = input(i)
		}

		if k[0] == k[1] {
			fmt.Println("您输入的两条线是平行的！")
		} else {
			fmt.Println("您输入的两条线不是平行的！")
		}

		var whetherContinue string
		fmt.Print("需要重新计算两条线是否平行吗？（y/n）:")
		fmt.Scanln(&whetherContinue)
		if whetherContinue != "y" {
			return
		}
	}
}

/*
	输入方法
	TODO 校验输入
*/
func input(i int) float64 {
	var coordX1 float64
	fmt.Print("请输入第", i+1, "条直线的第一个坐标（x）：")
	fmt.Scan(&coordX1)

	var coordY1 float64
	fmt.Print("请输入第", i+1, "条直线的第一个坐标（Y）：")
	fmt.Scan(&coordY1)

	var coordX2 float64
	fmt.Print("请输入第", i+1, "直线的第二个坐标（x）：")
	fmt.Scan(&coordX2)

	var coordY2 float64
	fmt.Print("请输入第", i+1, "条直线的第二个坐标（y）：")
	fmt.Scan(&coordY2)

	fmt.Println("")
	//计算两个斜率，并保存到k数组中
	return calc(coordX1, coordY1, coordX2, coordY2)

}

/*
	计算斜率
*/
func calc(coordX1 float64, coordY1 float64, coordX2 float64, coordY2 float64) float64 {
	//TODO 其他情况
	return (coordY2 - coordY1) / (coordX2 - coordX1)
}
