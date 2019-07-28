package main

import "fmt"

func main() {
	//数组参数切片
	num := [5]int{1, 2, 3, 4, 5}
	s := num[0:2]
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", num)

	fmt.Printf("%p %p\n", s, &num[0])
	//切片是指针，和数组指向同一个内存空间
	s[0] = 6
	fmt.Println(s, num)
	//切片添加
	s = append(s, 7)
	fmt.Println(s, num)
	//当切片添加元素大于数组，切片会重新开辟空间
	s = append(s, 8, 9, 10)
	fmt.Println(s, num)
	fmt.Printf("%p %p\n", s, &num[0])

	//删除代码
	n :=2
	newSlice :=s[0:n]
	newSlice = append(newSlice,s[n+1:]...)
	fmt.Println(s)
	fmt.Println(newSlice)
}
