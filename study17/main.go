package main

import (
	"container/ring"
	"fmt"
)

func main() {
	//ring,循环链表
	r := ring.New(5)
	r.Value = 0
	r.Next().Value = 1
	r.Next().Next().Value = 2
	r.Prev().Value = 4
	r.Prev().Prev().Value = 3

	r.Do(func(i interface{}) {
		fmt.Println(i)
	})
	//fmt.Println(r)
	fmt.Println("\n")
	//删除,取值n%r.len()
	r.Unlink(7)
	r.Do(func(i interface{}) {
		fmt.Println(i)
	})

}
