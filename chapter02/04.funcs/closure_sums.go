package main

import "fmt"

var counter int

func calcSum(nums ...int) (sum int) {
	for _, item := range nums {
		sum += item
	}
	counter++
	return
}

func showUseTimes() {
	fmt.Println(counter)
}

func genInprovementFunc() func(parentage float64) {
	base := 1000.0
	return func(parentage float64) {
		base += base * (1 + parentage)
		fmt.Println(base)
	}
}

func closuerMain() {
	imp := genInprovementFunc()
	imp(0.1)
	imp(0.1)
	imp(0.1)
	imp(0.1)

}
