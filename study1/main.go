package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world")
	fmt.Fprint(os.Stdout, "你好 世界")
	//返回字符串
	a := fmt.Sprint("hello world")
	fmt.Println(a)
}
