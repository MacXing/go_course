package main

import "fmt"

//封装
type People struct {
	name string
	age  int
}

func (p *People) SetName(name string) {
	p.name = name
}
func (p People) GetName() string{
	return p.name
}


func main() {
	peo := new(People)
	peo.SetName("allen")
	fmt.Printf(peo.GetName() )
}
