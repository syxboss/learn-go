package main

import (
	"fmt"
	"testing"
)

// 实现案例1：楼层有5层，所有电梯楼层没有人请求电梯，电梯不动
func TestCase1(t *testing.T) {
	elevator := Elevator{
		location: 1,
	}
	floor, err := elevator.move()
	if err != nil {
		t.Errorf("Elevator needs repair,reason:%s ", err)
	}
	fmt.Println("期望楼层：", floor)
}

//实现案例2：楼层有5层，电梯在1层。三楼按电梯。电梯向三楼行进，并停在三楼。
func TestCase2(t *testing.T) {
	elevator := Elevator{
		location:    1,
		targetFloor: []int{3},
	}
	floor, err := elevator.move()
	if err != nil {
		t.Errorf("Elevator needs repair,reason:%s ", err)
	}
	fmt.Println("当前楼层：", floor)
}

//实现案例3：楼层有5层，电梯在3层。上来一些人后，目标楼层： 4楼、2楼。电梯先向上到4楼，然后转头到2楼，最后停在2楼。
func TestCase3(t *testing.T) {
	elevator := Elevator{
		location:    3,
		targetFloor: []int{4, 2, 2},
	}
	floor, err := elevator.move()
	if err != nil {
		t.Errorf("Elevator needs repair,reason:%s ", err)
	}
	fmt.Println("当前楼层：", floor)
}

//实现案例4：楼层有5层，电梯在3层。上来一些人后，目标楼层： 4楼、5楼、2楼。电梯先向上到4楼，然后到5楼，之后转头到2楼，最后停在2楼。
func TestCase4(t *testing.T) {
	elevator := Elevator{
		location:    3,
		targetFloor: []int{4, 5, 2},
	}
	floor, err := elevator.move()
	if err != nil {
		t.Errorf("Elevator needs repair,reason:%s ", err)
	}
	fmt.Println("当前楼层：", floor)
}
