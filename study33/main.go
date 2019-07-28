package main

import (
	"fmt"
	"sync"
)

//锁,互斥锁
//Mutex 互斥锁锁
//RWLock 读写锁

var (
	num = 100
	wg  sync.WaitGroup
	m   sync.Mutex
)

func demo() {
	m.Lock()
	for i := 0; i < 10; i++ {
		num--
	}
	m.Unlock()
	wg.Done()
}
func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go demo()
	}
	wg.Wait()
	fmt.Println(num)
	fmt.Println("end")
}
