package calculator

import (
	"fmt"
)

func BMI(weightKG, heightM float64) (bmi float64, err error) {
	fmt.Println("以下是最新是修改的BMI方法")
	if weightKG < 0 {
		err = fmt.Errorf("weight cannot be negative")
		return
	}
	if heightM < 0 {
		err = fmt.Errorf("height cannot be negative")
		return
	}
	if heightM == 0 {
		err = fmt.Errorf("height cannot be zero !")
		return
	}
	bmi = weightKG / (heightM * heightM)
	fmt.Println("===================")
	return
}
