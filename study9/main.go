package main

import "fmt"

func main() {
	//go语言,唯一循环：for循环
	//for i := 0; i < 5; i++ {
	//	fmt.Println(i)
	//}
	//死循环
	//for ; ;  {
	//	fmt.Println("执行")
	//
	//}

	arr := [3]int{1, 2, 3}
	//for i := 0; i < len(arr); i++ {
	//	fmt.Println(arr[i])
	//}
	for i, n := range arr {
		fmt.Println(i, n)
	}

}
