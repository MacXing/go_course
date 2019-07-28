package main

import (
	"fmt"
	"sort"
)

func main() {
	//sort
	num := []int{1, 3, 2, 7, 6, 1}
	//sort.Ints(num)
	//fmt.Println(num)

	sort.Sort(sort.IntSlice(num)) //升序
	fmt.Println(num)

	sort.Sort(sort.Reverse(sort.IntSlice(num))) //降序
	fmt.Println(num)

	f :=[]float64{1.2,3.4,5.1,0.2}
	sort.Sort(sort.Float64Slice(f))
	fmt.Println(f)

	//string类型 排序
	s :=[]string{"","allen","python","你好"}
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
}
