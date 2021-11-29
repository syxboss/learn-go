package main

import "fmt"

func main() {
	var h, s float64 = 3, 4.13
	var sc = h * s //h * s :Invalid operation: h * s (mismatched types int and float64)
	fmt.Println(h * s)
	fmt.Println(sc)
}
