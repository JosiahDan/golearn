package main

import (
	"fmt"
	"time"
)

//向通道发送数据
func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "BeiJing"
	ch <- "Tokyo"
}

//从通道取出数据
func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s\n", input)
	}
}

func main() {
	//实例化通道
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	//主线程睡眠等待协程运行
	time.Sleep(1e9)
}
