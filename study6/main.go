package study6

import (
	"fmt"
	"time"
)

func main() {
	//t := time.Now()
	//fmt.Println(t)
	////纳秒
	//t1 := time.Unix(0, t.UnixNano())
	//fmt.Println(t1.String())
	//
	//t2 :=time.Date(2022,2,4,14,8,8,0,time.Local)
	//fmt.Println(t2.String())
	//
	//fmt.Println(t.Year())
	//fmt.Println(int(t.Month()))
	//fmt.Println(t.Day())
	//fmt.Println(t)
	//fmt.Println(t.Date())
	//fmt.Println(t.Hour())
	//fmt.Println(t.Clock())
	//fmt.Println(t.Nanosecond())

	//时间和String转换

	//t := time.Now()
	//s :=t.Format("2006/01/02 15:04:05")
	//fmt.Println(s)

	s:="2022-2-4 18:8:9"
	time.Parse("2006-1-2 15:4:5",s)
	fmt.Println(s)
}
