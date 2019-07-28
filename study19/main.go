package main

import "fmt"

//是否大写决定是否能挎包访问，名字和属性名皆是如此
type User struct {
	Name string
	Age  int
}

func main() {
	//结构体
	user1 := User{"allen", 23}
	user2 := User{Name: "ZHOU", Age: 23}
	fmt.Println(user1, user2)
	fmt.Println(user1.Name)
}
