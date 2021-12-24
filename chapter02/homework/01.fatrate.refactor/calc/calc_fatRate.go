package calc

import calculator "github.com/armstrongli/learn.go/chapter03/001.fatrate.refactor/calc"

func CalcFatRate(bmi float64, age int, sex string) (fatRate float64) {
	fatRate = calculator.CalcFatRate(bmi, age, sex)
	return
}
