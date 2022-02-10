package main

import (
	"fmt"
	"time"
)

func main(){
	data := make(map[int]int,10)
	for i := 1;i <= 10;i++{
		data[i] = i
	}

	for key,value := range data {
		go func(key,value int) {
			fmt.Println("k ->",key,"v ->",value)
		}(key,value)
	}
	time.Sleep(time.Second * 5)
}