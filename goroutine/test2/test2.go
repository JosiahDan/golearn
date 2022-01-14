package main

import (
	"fmt"
	"time"
)

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	ch := make(chan int)
	go suck(ch)
	go pump(ch)
	time.Sleep(1e9)
}
