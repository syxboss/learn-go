package main

import "fmt"

type truckLegend struct {
}

func (*truckLegend) OpenTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 truckLegend 做OpenTheDoorOfRefrigerator")
	return nil
}
func (*truckLegend) PutElephantIntoRefrigerator(Elephant, Refrigerator) error {
	fmt.Println("用 truckLegend 做PutElephantIntoRefrigerator")
	return nil
}
func (*truckLegend) CloseTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 truckLegend 做CloseTheDoorOfRefrigerator")
	return nil
}
