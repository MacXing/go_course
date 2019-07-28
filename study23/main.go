package main

import "fmt"

//通过组合实现继承
type People struct {
	name string
	age  int
}
type Teacher struct {
	classroom string
	peo People
}

func main() {
	tea := Teacher{"abc", People{"allen", 23}}
	fmt.Printf(tea.classroom, tea.peo.name, tea.peo.age)
}
