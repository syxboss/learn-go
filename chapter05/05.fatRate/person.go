package main

import (
	gobmi "github.com/syxboss/awesomeProject"
	"log"
)

type Person struct {
	name   string
	sex    string
	tall   float64
	weight float64
	age    int

	bmi     float64
	fatRate float64
}

func (person *Person) calcBMI() error {
	bmi, err := gobmi.BMI(person.weight, person.tall)
	if err != nil {
		log.Printf("error when calculating BMP for Person[%s]:%v", person.name, person.bmi)
		return err
	}
	person.bmi = bmi
	return nil
}

func (person *Person) calcFatRate() {
	person.fatRate = gobmi.CalcFatRate(person.bmi, person.age, person.sex)
}
