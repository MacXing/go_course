package main

import "fmt"

func main() {
	var name,age string
	fmt.Println("请输入你的年龄和名字")
	fmt.Scan(&age,&name)
	fmt.Println("输入的内容:",name,age)
}
