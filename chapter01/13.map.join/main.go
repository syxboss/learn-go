package main

import "fmt"

func main() {
	// map addall 另一个map
	leftMap, rightMap := map[string]int{}, map[string]int{}

	leftMap["语文"] = 89
	rightMap["数学"] = 43

	for k, v := range rightMap {
		leftMap[k] = v
	}
	fmt.Println(leftMap)
}
