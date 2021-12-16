package main

import (
	"fmt"
	"os"
	"time"
)

func deferGuess() {
	startTime := time.Now()
	defer fmt.Println("duration:", time.Now().Sub(startTime)) //defer只管当前的函数，其他部分是准备好的，所以最终显示的时间并不是5s
	defer func() {
		fmt.Println("duration:", time.Now().Sub(startTime))
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("start time:", startTime)
}

/**
存在疑问，源码解答
*/
func openfile() {
	fileName := "/"
	fileObj, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	defer fileObj.Close()
	/**
	func (f *File) Close() error {
		if f == nil {
			return ErrInvalid // 为空处理
		}
		return f.file.close()
	}
	*/

	defer func() {
		fileObj.Close()
	}()
	fileObj = nil // 报错????
}
