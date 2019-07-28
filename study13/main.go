package main

import "fmt"

func main() {
	//copy,不会影响复制对象，严格脚标复制
	//s1 := []int{1, 2}
	//s2 := []int{3, 4, 5, 6}
	//copy(s2, s1)
	//fmt.Println(s1, s2)

	//cpoy删除元素
	s := []int{1, 2, 3, 4, 5, 6}
	n := 2
	newSlice := make([]int, n)
	copy(newSlice, s[0:n])
	newSlice = append(newSlice, s[n+1:]...)
	fmt.Println(newSlice)
	fmt.Println(s)

}
