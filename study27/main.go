package main

import (
	"fmt"
	"os"
)

func main() {
	fp := "./resource/1.txt"
	read(fp)
	write(fp, "测试\n")
	read(fp)
}

func write(fp string, sentence string) {
	f, err := os.OpenFile(fp, os.O_APPEND, 0660)
	defer f.Close()
	if err != nil {
		f, _ = os.Create(fp)
	}
	_, _ = f.Write([]byte(sentence))
}
func read(path string) {
	//Open 只读
	file, _ := os.Open(path)
	fileInfo, _ := file.Stat()
	b := make([]byte, fileInfo.Size())
	_, _ = file.Read(b)
	fmt.Println("文件内容为：", string(b))
}
