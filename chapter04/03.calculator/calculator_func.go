package main

import "fmt"

// 能用指针尽量用指针
func (c Calculator) Add() int { // (c *Calculator)传入指针，指针在运行过程中可以修改值
	fmt.Printf("add &c=%p \n", &c)
	tempResult := c.left + c.right
	c.result = tempResult
	fmt.Println("调用add函数，result=", c.result)
	return tempResult
}

func (c Calculator) Sub() int {
	return c.left - c.right
}

func (c Calculator) Multiple() int {
	return c.left * c.right
}

func (c Calculator) Divide() int {
	return c.left / c.right
}

func (c Calculator) Reminder() int {
	return c.left % c.right
}

//只当执行该方法的时候会修改result的值
func (c *Calculator) SetResult(result int) {
	c.result = result
}
