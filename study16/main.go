package main

import (
	"container/list"
	"fmt"
)

func main() {
	//list,go语言对双向链表的实现
	myList := list.New()
	fmt.Println(myList)
	fmt.Printf("%T\n", myList)

	myList.PushFront("a")
	myList.PushBack("z")
	myList.InsertAfter("b", myList.Back())
	fmt.Println(myList)

	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	//fmt.Println(myList.Front().Value)
	//fmt.Println(myList.Front().Next().Value)
	//
	//n := 3
	//for i := 0; i < n; i++ {
	//	myList.Front().Next()
	//}
	//var curr *list.Element
	//if n > 0 && n <= myList.Len() {
	//	if n == 1 {
	//		curr = myList.Front()
	//	} else if n == myList.Len() {
	//		curr = myList.Back()
	//	} else {
	//		curr = myList.Front()
	//		for i := 0; i < n; i++ {
	//			curr = curr.Next()
	//		}
	//	}
	//	fmt.Println("取出：", curr.Value)
	//} else {
	//	fmt.Println("数字不对")
	//}
	fmt.Print("\n")
	myList.MoveAfter(myList.Front(), myList.Back())
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
}
