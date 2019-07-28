package main

import "fmt"

func main() {
	//切片,具有可变长度的相同元素序列，切片声明不会开辟内存,是引用类型
	//var slice []string //切片
	//var arr [3]string  //数组
	//fmt.Println(slice,arr)
	//fmt.Println(slice==nil)
	//fmt.Printf("%p",&slice)
	name := []string{"allen", "python", "go"}
	fmt.Println(name)
	fmt.Printf(name[0])
	//fmt.Printf(name[0:3])



}
