package main

import (
	"fmt"
)

func makeNum(numCh chan int) {
	for i := 0; i < 2000; i++ {
		numCh <- i
	}
}

func getNUm(numCh chan int, resCh chan int) {
	for {
		v, ok := <-numCh
		if !ok {
			break
		}

		resCh <- v
	}
}

func main() {
	numCh := make(chan int)
	resCh := make(chan int)
	go makeNum(numCh)

	for i := 0; i < 8; i++ {
		go getNUm(numCh, resCh)
	}

	for v := range resCh {
		fmt.Println(v)
	}
}
