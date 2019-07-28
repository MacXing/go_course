package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//通过纳秒差 随机种子
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(10)
	fmt.Println(a)
}
