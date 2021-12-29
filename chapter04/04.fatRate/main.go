package main

import "fmt"

func main() {
	person := getFakePersonInfo()
	c := Calc{}
	c.BMI(person)
	c.fatRate(person)
	fmt.Println(person)
	sug := suggestion{}
	suggestions := sug.GetSuggestion(person)
	fmt.Println(suggestions)
}

func getMaterialsFromImput() *Person {
	// 录入各项
	var name string
	fmt.Print("请输入你的名字:")
	fmt.Scanln(&name)

	var weight float64
	fmt.Print("请输入体重【kg】:")
	fmt.Scanln(&weight) //取地址

	var tall float64
	fmt.Print("请输入身高【m】:")
	fmt.Scanln(&tall)

	var age int
	fmt.Print("请输入年龄【岁】:")
	fmt.Scanln(&age)

	var sex string
	fmt.Print("请输入sex【男、女】:")
	fmt.Scanln(&sex)

	return &Person{
		name:   name,
		sex:    sex,
		tall:   tall,
		weight: weight,
		age:    age,
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
