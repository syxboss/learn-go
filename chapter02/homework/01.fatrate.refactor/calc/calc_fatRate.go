package calc

import "fmt"

func CalcFatRate(bmi float64, age int, sex string) (fatRate float64, err error) {
	if age <= 0 {
		err = fmt.Errorf("Your age is less than or equal to 0. What are you?")
		return
	} else if age < 18 {
		err = fmt.Errorf("Sorry ，minors will not be involved in the calculation !")
		return
	} else if age > 150 {
		err = fmt.Errorf("You are older than the maximum age limit of 150 years!")
	}

	if bmi <= 0 {
		err = fmt.Errorf("BMI cannot be less than or equal to 0!")
		return
	}

	sexWeight := 0
	switch sex {
	case "男":
		sexWeight = 1
	case "女":
		sexWeight = 0
	default:
		err = fmt.Errorf("Your gender is unknown, " +
			"currently only male and female can be identified, please specify your gender !")
	}

	fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	return
}
