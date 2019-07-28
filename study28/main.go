package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//fp := "./resource/1.txt"
	//f, _ := os.Open(fp)
	//b, _ := ioutil.ReadAll(f)
	//fmt.Println(string(b))
	//b, _ := ioutil.ReadFile(fp)
	//fmt.Println(string(b))
	//ioutil.WriteFile(fp, []byte("hello world \n"), 0666)
	fi, _ := ioutil.ReadDir("./resource")
	for _, n := range fi {
		fmt.Println(n.Name())
	}
}
