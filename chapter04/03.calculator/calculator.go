package main

var defaultCalculator = Calculator{}

//跨包私有变量不可见
type Calculator struct {
	left, right int
	op          string

	result int
}
