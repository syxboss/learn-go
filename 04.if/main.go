package main

import "fmt"

func main() {
	var fruit string = "6 apple"
	var watermallan bool = false // true and false
	if watermallan {
		fruit = "1个apple"
	} else {
		fruit = "7个apple"
	}
	fmt.Println(fruit)
}
