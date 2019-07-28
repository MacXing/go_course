package main

import "fmt"

func main() {
	/*
		算法运算
	*/
	a, b, c := 1, 2, 3
	fmt.Println(a + b*c)
	a++
	fmt.Println(a)
	b--
	fmt.Println(b)
	//取余判读数字是否可以被整除
	fmt.Println(b % c)
}
