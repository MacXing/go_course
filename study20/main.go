package main

import "fmt"

type Uer struct {
	Name string
	Age  int
}

func main() {
	//结构体指针
	user1 := new(Uer)
	user2 := &Uer{"allen", 21}
	fmt.Printf("%p %p\n", user1, user2)
}
