package main

import "fmt"

func main() {
	names := []string{"小强", "周毅", "旭东"}
	fr := []float64{28, 18, 18}

	names = append(names, "张松")
	fr = append(fr, 19)

	for i, name := range names {
		if name == "周毅" {
			fmt.Println("%s的体脂率时候 %f\n", name, fr[i])
		}
	}

	// 定义
	var m1 map[string]int
	m2 := map[string]int{}
	m3 := map[string]int{"王强": 60, "李静": 83, "苗文": 91}
	fmt.Println(m1, m2, m3)

	mSurprise := map[string]map[string]map[struct{}]int{}
	fmt.Println(mSurprise)

	// 增删改查
	m3["小强"] = 77
	fmt.Println(m3)
	delete(m3, "小强")
	fmt.Println(m3)
	m3["小强"] = 80
	fmt.Println(m3)
	fmt.Println(m3["小强"])
	fmt.Println(m3["张三"]) // 假值：0

	xqScore, ok := m3["张三"] //xqScore 真或假, ok optional 可选返回值，是否在map中
	fmt.Println(xqScore, ">>>>>", ok)

	// 遍历
	for k, v := range m3 {
		fmt.Println(k, " ", v)
	}

	var m4 map[string]int = nil
	//m4["a"] =1 //panic: assignment to entry in nil map
	delete(m4, "a")
	fmt.Println("m1没有实例化，直接取数：", m4["a"])
}
