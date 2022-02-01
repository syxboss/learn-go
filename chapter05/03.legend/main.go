package main

func main() {
	//fmt.Println("传说中，装大象可以这么装：")
	//
	//var legend PutElephantIntoRefrigerator = &manPushLegend{}
	//var r Refrigerator
	//var e Elephant
	//
	//legend.OpenTheDoorOfRefrigerator(r)
	//legend.PutElephantIntoRefrigerator(e,r)
	//legend.CloseTheDoorOfRefrigerator(r)
	//
	//fmt.Println("this is a legendary")
	legendary(&manPushLegend{}, Refrigerator{}, Elephant{})
}
