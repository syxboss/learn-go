package calc

import (
	"fmt"
	"github.com/syxboss/awesomeProject"
)

func CalcBMI(tall float64, weight float64) (bmi float64, err error) {
	bmi, err = gobmi.BMI(weight, tall)
	fmt.Println("以下是原来计算BMI的方法")
	if tall <= 0 {
		panic("身高不能为0或者负数")
	}
	//bmi = weight / (tall * tall)
	return
}
