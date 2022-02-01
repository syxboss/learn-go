package main

import "fmt"

func main() {
	//person := getFakePersonInfo()
	//c := Calc{}
	//c.BMI(person)
	//c.fatRate(person)
	//fmt.Println(person)
	//sug :=fatRateSuggestion{}
	//suggestions := sug.GetSuggestion(person)
	//fmt.Println(suggestions)

	frSvc := &fatRateService{
		s: GetFatRateSuggestion(),
	}
	fakePerson := getFakePersonInfo()
	fmt.Println(frSvc.s.GiveSuggestionPerson(fakePerson))

	for {
		p := InputService().GetInput()
		fmt.Println(frSvc.s.GiveSuggestionPerson(p))
	}

}

func getFakePersonInfo() *Person {
	return &Person{
		name:   "小强",
		sex:    "男",
		tall:   1.7,
		weight: 70,
		age:    35,
	}
}
