package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

//打开一个文件并读取其内容
func openAndRead() {
	//打开文件并判断文件是否打开错误
	inputFile, inputError := os.Open("./records/dat1.txt")

	if inputError != nil {
		fmt.Println("打开文件失败")
		return
	}

	//如果文件打开没有错误，最后要关闭文件
	defer inputFile.Close()

	//创建读取器
	inputReader := bufio.NewReader(inputFile)
	//通过'\n'进行行读取
	for {
		//如果读取完文件readerError值会变为io.EOF退出文件
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("文件输入为:%s", inputString)
		//io.EOF为io包中定义的错误类型,意为文件中没有输入了,可用来判断是否完全读取文件
		if readerError == io.EOF {
			return
		}
	}
}

func openAndWrite() {
	outputFile, outputError := os.OpenFile("./records/dat1.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Println("打开文件失败")
	}

	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := time.Now().String() + " hello,world\n"

	outputWriter.WriteString(outputString)
	outputWriter.Flush()
	fmt.Println("写入成功")
}

func main() {
	openAndWrite()
}
