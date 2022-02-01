package main

import "log"

type fatRateService struct {
	s *fatRateSuggestion
}

func (frSvc *fatRateSuggestion) GiveSuggestionPerson(person *Person) string {
	if err := person.calcBMI(); err != nil {
		log.Printf("无法给%s计算BMI: 原因是：%v", person.name, err)
	}
	person.calcFatRate()
	return frSvc.GetSuggestion(person)
}

func (frSvc *fatRateSuggestion) GiveSuggestionPersons(persons ...*Person) map[*Person]string {
	out := map[*Person]string{}
	for _, item := range persons {
		out[item] = frSvc.GiveSuggestionPerson(item)
	}
	return out
}
