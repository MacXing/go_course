package main

import (
	"fmt"
	"unsafe"
)

/*
	变量：所有东西都是变量

*/
func main() {
	//声明变量

	//var name = "allen"
	//fmt.Println(name)
	////打印类型
	//fmt.Printf("%T",name)

	//短变量,只能在函数内
	name := "allen"
	fmt.Println(name)

	var (
		a = 1
		b = 2
		c = "allen"
	)
	fmt.Println(a, b, c)
	//占用字节
	d := unsafe.Sizeof(a)
	fmt.Println(d)
}
