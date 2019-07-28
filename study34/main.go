package main

import "fmt"

//channel,线程通信
//func main() {
//	ch := make(chan int)
//	go func() {
//		fmt.Println("执行")
//		fmt.Println("结束")
//		ch <- 998 //放在最后一行
//	}()
//	a := <-ch
//	fmt.Println(a)
//	fmt.Println("结束")
//}
/*
	两个协程通信
*/
//func main() {
//	ch1 := make(chan string)
//	ch2 := make(chan int)
//	go func() {
//		ch1 <- "传送宁一个协程的数据"
//		ch2 <- 1
//	}()
//	go func() {
//		content := <-ch1
//		fmt.Println("取出数据：", content)
//		ch2 <- 2
//	}()
//
//	<-ch2
//	<-ch2
//	fmt.Println("结束")
//}
/*
channel是安全的
*/
func main() {
	ch1 := make(chan int)
	for i := 0; i < 10; i++ {
		go func(j int) {
			fmt.Println("开始写")
			ch1 <- j
			fmt.Println("结束写")
		}(i)
	}

	for j := 0; j < 10; j++ {
		a := <-ch1
		fmt.Println(a)
	}
}
