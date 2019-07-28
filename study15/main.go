package main

import "fmt"

func main() {
	//map
	//var m map[string]string
	//fmt.Println(m)
	//fmt.Printf("%p",m)

	//m := make(map[string]string)
	//fmt.Println(m == nil)
	//fmt.Printf("%p", m)

	m := map[string]string{"key": "value", "name": "allen"}
	fmt.Println(m["key"], m["name"], m["python"])
	m["python"] = "go"
	fmt.Println(m["key"], m["name"], m["python"])

	//删除
	delete(m, "name")
	fmt.Println(m)

	for key, value := range m {
		fmt.Println(key, value)
	}
}
