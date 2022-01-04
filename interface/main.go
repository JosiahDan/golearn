package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

//使用接口的方式实现一个既可以往终端写日志也可以往文件写日志的简易日志库。

type fileRecord struct {
}

func (f fileRecord) write() {
	//打开文件如果不存在则创建文件
	inputFile, inputError := os.OpenFile("./records/dateRecords.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if inputError != nil {
		fmt.Println("文件打开失败")
		return
	}
	//程序结束后关闭文件
	defer inputFile.Close()

	inputWriter := bufio.NewWriter(inputFile)
	inputString := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05") + " hello,world\n"

	inputWriter.WriteString(inputString)
	inputWriter.Flush()
	fmt.Println("写入完成")
}

type logRecord struct {
}

func (l logRecord) write() {
	fmt.Println(time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05") + " hello,world\n")
}

type recorder interface {
	write()
}

func main() {
	var file fileRecord
	var log logRecord
	var x recorder

	x = file
	x.write()
	x = log
	x.write()
}
