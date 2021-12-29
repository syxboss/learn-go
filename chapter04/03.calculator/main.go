package main

import (
	"fmt"
)

func main() {
	var left, right int = 1, 2
	var op string = "+"
	c := Calculator{
		left:  left,
		right: right,
		op:    op,
	}
	fmt.Printf("main &c=%p \n", &c)
	fmt.Println(c.Add())
	fmt.Println("c.result=", c.result) //装不进去，c.result= 0

	//newC := Newcalculator{&Calculator{}}
	newC := Newcalculator{}
	newC.left = left
	newC.right = right
	fmt.Println(newC.Add())

	mc := MyCommand{}
	mc.commandOption["aa"] = "AAA"
	fmt.Println(mc.ToCmdStr())

	//fmt.Scanln(&left)
	//fmt.Print(" ")
	//fmt.Scanln(&op)
	//fmt.Print(" ")
	//fmt.Scanln(&right)
	//
	//switch op {
	//case "+":
	//	fmt.Println(left + right)
	//case "-":
	//	fmt.Println(left - right)
	//case "*":
	//	fmt.Println(left * right)
	//case "/":
	//	fmt.Println(left / right)
	//case "%":
	//	fmt.Println(left % right)
	//default:
	//	fmt.Println("not supported calculator rule")
	//}
}

type MyCommand struct {
	commandOption map[string]string
}

func (my MyCommand) ToCmdStr() string {
	out := ""
	for k, v := range my.commandOption {
		out += fmt.Sprintf("---%s==%s---", k, v)
	}
	return out
}
