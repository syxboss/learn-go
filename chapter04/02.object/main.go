package main

import (
	"fmt"
	gobmi "github.com/syxboss/awesomeProject"
)

// 继承有点特殊，现在不讲
type Person struct {
	name   string
	sex    string
	tall   float64
	weight float64
	age    int
}

type human struct{}

func main() {
	//name1 ,sex1 , tall1 , weight1 ,sex1 : = // ....
	//name2 ,sex2 , tall2 , weight2 ,sex2 : = // ....
	//name3 ,sex3 , tall3 , weight3 ,sex3 : = // ....
	//
	//names := []string{}
	//sexs :=[]string{}
	//talls := []float64{}
	//weight := []float64{}
	//ages := []int{}

	persons := []Person{
		{
			"小强",
			"男",
			1.7,
			70,
			35,
		},
	}

	//正确的实现方式
	xq := Person{
		name:   "小强",
		sex:    "男",
		tall:   1.7,
		weight: 70,
		age:    35,
	}
	fmt.Println(xq)

	a := new(Person) // *Person
	fmt.Println(a)

	for _, item := range persons {
		bmi, err := gobmi.BMI(item.weight, item.tall)
		fmt.Println(bmi, err)
	}
}
