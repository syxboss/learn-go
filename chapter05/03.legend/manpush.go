package main

import "fmt"

type manPushLegend struct {
}

func (*manPushLegend) OpenTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 manPushLegend 做OpenTheDoorOfRefrigerator")
	return nil
}
func (*manPushLegend) PutElephantIntoRefrigerator(Elephant, Refrigerator) error {
	fmt.Println("用 manPushLegend 做PutElephantIntoRefrigerator")
	return nil
}
func (*manPushLegend) CloseTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 manPushLegend 做CloseTheDoorOfRefrigerator")
	return nil
}
