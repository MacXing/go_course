package main

import (
	"fmt"
	"os"
)

func main() {
	createDir("resource")
	createFile("./resource/1.txt")
	//os.Rename()

}

func createDir(dir string) {
	error := os.Mkdir(dir, os.ModeDir)
	if error != nil {
		fmt.Println("创建失败", error)
		return
	}
	fmt.Println("创建成功！")

}
func createFile(name string) {
	file, error := os.Create(name)
	if error != nil {
		fmt.Println("文件创建失败")
		return
	}
	fmt.Println("文件创建成功！", file)
}
