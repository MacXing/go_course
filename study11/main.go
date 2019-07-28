package main

import "fmt"

func main() {
	//make函数，slice,map,channel,interface,使用make函数可以声明，但又不是nil
	//s := make([]int, 0)
	//fmt.Println(s)
	//fmt.Printf("%p\n", s)
	////长度表示切片中元素个数，cap容量表示切片空间大小
	//fmt.Println(len(s), cap(s))

	//append,直接添加到切片中,如果切片长度大于容量，容量翻倍
	s := make([]string, 0)
	fmt.Println(len(s), cap(s))
	s = append(s, "allen")
	fmt.Println(s)
	fmt.Println(len(s),cap(s))
	s = append(s, "go")
	fmt.Println(s)
	fmt.Println(len(s),cap(s))
	s = append(s, "python")
	fmt.Println(s)
	fmt.Println(len(s),cap(s))
	s = append(s, "pyzhou")
	fmt.Println(s)
	fmt.Println(len(s),cap(s))
	s = append(s, "gozhou")
	fmt.Println(s)
	fmt.Println(len(s),cap(s))
	s = append(s, "a","b","c","d")
	fmt.Println(s)
	fmt.Println(len(s),cap(s))

}
