package main

import (
	"fmt"
	"time"
)

//线程，延时执行,goroutine
//go 表达式
func main() {
	//fmt.Println("程序开始")
	//time.AfterFunc(3e9, func() {
	//	fmt.Println("延时执行")
	//})
	//time.Sleep(4e9)
	//fmt.Println("程序结束")
	for i := 0; i < 5; i++ {
		go demo(i)
	}
	time.Sleep(3e9)
}

func demo(count int) {
	for i := 0; i < 100; i++ {
		fmt.Println(count, i)
	}
}
