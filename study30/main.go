package main

import (
	"log"
	"os"
)

//日志
func main() {
	//普通日志
	//log.Print("打印日志信息")
	////错误信息
	//log.Panicln("打印panic信息")
	////非常严重
	//log.Fatal("打印fatal信息")
	f, _ := os.OpenFile("./resource/golog.log", os.O_APPEND|os.O_CREATE, 0777)
	logger := log.New(f, "Info:", log.Ltime)
	logger.Println("打印日志信息")
}
