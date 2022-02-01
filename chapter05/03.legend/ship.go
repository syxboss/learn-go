package main

import "fmt"

type shipLegend struct {
}

func (*shipLegend) OpenTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 shipLegend 做OpenTheDoorOfRefrigerator")
	return nil
}
func (*shipLegend) PutElephantIntoRefrigerator(Elephant, Refrigerator) error {
	fmt.Println("用 shipLegend 做PutElephantIntoRefrigerator")
	return nil
}
func (*shipLegend) CloseTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 shipLegend 做CloseTheDoorOfRefrigerator")
	return nil
}
