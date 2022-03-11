package main

import (
	"fmt"
	"os"
)

type Persons []*Person

type Person struct {
	name    string
	fatRate float64
	//weight float64
	//height float64
	//...
}

type FatRateScope struct {
	min float64 // 最小值
	max float64 // 最大值
}

//person -> json
type PersonJson struct {
	Name    string  `json:"name" `
	FatRate float64 `json:"fatrate"`
}

// 注册内容,可以在这里计算体脂（输入体重等信息）
func (p *Person) register(name string, fatRate float64) {
	p.name = name
	p.fatRate = fatRate //calcFatRate()
}

// 文件增加JSON数据，一行一个人的数据
func (p *Person) appendPersonIntoFile(data []byte, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	defer file.Close()
	if err != nil {
		fmt.Println("无法打开", filePath, "下的文件：", file.Name(), "错误信息是：", err)
		os.Exit(1)
	}
	_, err = file.Write(append(data, '\n'))
	if err != nil {
		return fmt.Errorf("无法追加回车，错误信息是：", err)
	}
	return nil
}
