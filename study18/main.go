package main

import (
	"fmt"
)

func main() {
	//count := add(1, 2)
	//fmt.Println(count)
	//display("测试")
	//a, b := mutiBck()
	//fmt.Println(a, b)
	//demo("allen", "python", "go")
	//匿名函数
	//func() {
	//	fmt.Println("匿名函数")
	//}()
	//
	//func(name string) {
	//	fmt.Println(name)
	//}("go")
	//
	//f := func(name string) (string) {
	//	fmt.Println(name)
	//	return "go"
	//}
	//fmt.Println(f("python"))
	//mydo(func(name string) {
	//	fmt.Println(name)
	//})
	fmt.Println(a()())
}
func display(s string) {
	fmt.Println(s)
}
func add(a int, b int) int {
	return a + b
}

func mutiBck() (int, int) {
	fmt.Println("多返回函数")
	return 1, 2
}

//可变参数函数
func demo(hover ...string) {
	for i, n := range hover {
		fmt.Println(i, n)
	}
}

//引用类型，go语言中的五种：函数变量，slice，map,channel,interface
func mydo(arg func(name string)) {
	fmt.Println("执行mydo")
	arg("python")
}

func a() func() int {
	return func() int {
		return 100
	}
}
