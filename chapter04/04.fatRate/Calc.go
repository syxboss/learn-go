package main

import (
	gobmi "github.com/syxboss/awesomeProject"
	"log"
)

type Calc struct {
}

func (Calc) BMI(person *Person) error {
	bmi, err := gobmi.BMI(person.weight, person.tall)
	if err != nil {
		// 出现错误时，日志最好贴近错误
		log.Println("error when calculation bmi :", err)
		return err
	}
	person.bmi = bmi
	return nil
}

func (Calc) fatRate(person *Person) error {
	person.fatRate = gobmi.CalcFatRate(person.bmi, person.age, person.sex)
	return nil
}
