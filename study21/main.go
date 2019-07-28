package main

import "fmt"

type People struct {
	Name string
	Age  int
}

//方法，修改里面内容，用指针类型
func (p People) run() {
	fmt.Println(p.Name, "正在跑步，年龄为：", p.Age)
}
func main() {
	//方法
	peo := People{"allen", 23}
	peo.run()
}
