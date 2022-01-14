package main

import "fmt"

const (
	top = 5
	low = 1
)

type Elevator struct {
	location    int
	direction   string
	targetFloor []int
}

func (e *Elevator) move() (currentFloor int, err error) {
	//currentFloor = e.location
	if e.location < 1 || e.location > 5 {
		err = fmt.Errorf("电梯楼层数错误！您所按的楼层：%v", e.location)
		return
	}

	if len(e.targetFloor) > 0 {

		for _, item := range e.targetFloor {
			floor := item
			fmt.Printf("当前楼层：%v ,目标楼层：%v\n", e.location, floor)
			e.determineDirection(floor)
			if e.direction == "上" {
				fmt.Printf("电梯向上移动，从 %v 层到 %v 层\n", e.location, item)
				e.location = floor
				e.direction = ""
			} else if e.direction == "下" {
				fmt.Printf("电梯向下移动，从 %v 层到 %v 层\n", e.location, item)
				e.location = floor
				e.direction = ""
			}
		}
	}
	//e.determineDirection(e.location)

	if e.direction == "" {
		currentFloor = e.location
		fmt.Printf("电梯停止在当前楼层：%v \n", e.location)
		return
	}
	currentFloor = e.location
	err = fmt.Errorf("")
	return
}

func (e *Elevator) determineDirection(currentFloor int) {
	if currentFloor > e.location {
		e.direction = "上"
	} else if currentFloor < e.location {
		e.direction = "下"
	}
}
