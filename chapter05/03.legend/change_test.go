package main

import (
	"fmt"
	"testing"
)

type Change interface {
	ChangeName(newName string)
	ChangeAge(newAge int)
}

type Student struct {
	Name string
	Age  int
}

func (s *Student) ChangeName(newName string) {
	s.Name = newName
}

func (s *Student) ChangeAge(newAge int) {
	s.Age = newAge
}

func TestVal(t *testing.T) {

	//Cannot use 'Student{}' (type Student) as the type
	//Change Type does not implement 'Change' as the 'ChangeName' method has a pointer receiver
	//stdChg = Student{}
	var stdChg Change
	stdChg = &Student{
		Name: "Tom",
		Age:  23,
	} // 正确
	fmt.Println(stdChg)
}

func TestStd(t *testing.T) {
	s := Student{Name: "Tom"}
	s.ChangeName("Jerry")
	fmt.Println(s.Name)
}
