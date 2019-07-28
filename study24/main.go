package main

import "fmt"

//defer,完成延迟功能，长用于关闭文件、数据库连接等
//多个defer，多重defer采用栈结构，先产生后执行
func main() {
	//fmt.Println("打开连接")
	//defer func() {
	//	fmt.Println("关闭连接")
	//}()
	//fmt.Println("执行功能")
	fmt.Println("打开A")
	defer fmt.Println("关闭A")
	fmt.Println("打开B")
	defer fmt.Println("关闭B")
}
