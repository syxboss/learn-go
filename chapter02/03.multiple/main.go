package main

import "fmt"

func main() {
	bmis := []float64{1, 2, 3, 4.323, 32.4214}
	avgBmi := caculateAvg(bmis...)      // ...不定长
	avgBmi1 := caculateAvgOnSlice(bmis) // ...不定长
	fmt.Println(avgBmi)
	fmt.Println(avgBmi1)

	getScoresOfStudents("Tom")                       // func getScoresOfStudents(stdName string) (int, int, int, int, int)
	i1, i2, i3, i4, i5 := getScoresOfStudents("tom") // func getScoresOfStudents(stdName string) (chinese int, math int, english int, geo int, music int)
	fmt.Println(i1, i2, i3, i4, i5)
}

func caculateAvgOnSlice(bmis []float64) float64 {
	var total float64 = 0 // total := 0
	for _, item := range bmis {
		total += item //Invalid operation: total+=item (mismatched types int and float64)
	}
	return total / float64(len(bmis))
}

func caculateAvg(bmis ...float64) (avgBmi float64) {
	var total float64 = 0 // total := 0
	for _, item := range bmis {
		total += item //Invalid operation: total+=item (mismatched types int and float64)
	}
	avgBmi = total / float64(len(bmis))
	return
}

func getScoresOfStudents(stdName string) (chinese int, math int, english int, geo int, music int) {
	return 0, 0, 0, 0, 0
}
