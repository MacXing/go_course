package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 1.5
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))

}
