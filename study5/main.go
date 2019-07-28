package main

import "fmt"

func main() {
	//if score:=70;score>=60{
	//	fmt.Println("及格")
	//}else {
	//	fmt.Println("不及格")
	//}
	score := 81
	if score > 60 && score < 80 {
		fmt.Println("及格")
	} else if score < 60 {
		fmt.Println("不及格")
	} else if score > 80 {
		fmt.Println("优秀")
	}

	num := 10
	switch num {
	case 2:
		fmt.Println("二进制")
	case 8:
		fmt.Println("八进制")
	case 10:
		fmt.Println("十进制")
	default:
		fmt.Println("没有这个进制")
	}
}
