package main

import "fmt"

// 强制实现接口所有的方法，没有实现就报错
// _ 忽略值
var _ Door = &GlassDoor{}

type GlassDoor struct {
}

func (d *GlassDoor) unlock() {
	fmt.Println("GlassDoor unlock")
}

func (d *GlassDoor) lock() {
	fmt.Println("GlassDoor lock")
}

func (*GlassDoor) Open() {
	fmt.Println("GlassDoor Open")
}
func (*GlassDoor) Close() {
	fmt.Println("GlassDoor Close")
}
