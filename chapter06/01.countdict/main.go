package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	fmt.Println("开始计数")
	var totalCount int64 = 0
	for i := 0; i <= 5000; i++ {
		fmt.Println("正在统计第", i, "页")
		time.Sleep(1 * time.Second)
		r, _ := rand.Int(rand.Reader, big.NewInt(800))
		fmt.Println("有", r.Int64(), "字")
		totalCount += r.Int64()
	}
	fmt.Println("总共有：", totalCount, "字")
}
