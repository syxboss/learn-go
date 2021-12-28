package go_bmi

import "fmt"

func BMI(weight, height float64) (bmi float64, err error) {
	fmt.Println("以下是在本地扩展过的BMI方法")
	if weight <= 0 {
		err = fmt.Errorf("weight cannot be negative or zero!")
		return
	}
	if height <= 0 {
		err = fmt.Errorf("height cannot be negative or zero!")
		return
	}
	fmt.Println("计算BMI")
	bmi = weight / (height * height)
	fmt.Println("计算BMI结束，BMI：", bmi)
	return
}
