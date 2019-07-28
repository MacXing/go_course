package main

import "fmt"

func main() {
	//数组,数组在go语言中，是值类型
	var arr1 [3] int = [3] int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3}
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(len(arr1), len(arr2), len(arr3))
	arr1[1] = 10
	fmt.Println(arr1[1])
	//值类型复制，重新开辟一个空间复制
	arr4 := arr1
	fmt.Printf("%p,%p", &arr1, &arr4)

}
