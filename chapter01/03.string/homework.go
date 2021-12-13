package main

import (
	"fmt"
	"math"
)

/**
计算两个圆的面积之和并输出，精确到小数点后3位
提示：area=πr^2  %.3f 默认宽度，精度3
*/
func main() {
	r1 := 3.0
	area1 := math.Pi * math.Pow(r1, 2)
	fmt.Printf("第一个圆的面积：%.3f\n", area1)

	r2 := 3.45
	area2 := math.Pi * math.Pow(r2, 2)
	fmt.Printf("第二个圆的面积：%.3f\n", area2)

	fmt.Printf("圆的面积之和：%.3f\n", area1+area2)
}
