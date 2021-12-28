package calc

import (
	go_bmi "github.com/armstrongli/go-bmi"
)

func CalcBMI(tall float64, weight float64) (bmi float64, err error) {
	bmi, err = go_bmi.BMI(weight, tall)
	return
}
