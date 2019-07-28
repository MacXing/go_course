package main

import "fmt"

//select
func main() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func(j int) {
			ch <- j
		}(i)
	}

	for i := 0; i < 10; i++ {
		select {
		case a := <-ch:
			fmt.Println(a)
		default:
			
		}

	}
}
