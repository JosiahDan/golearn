package main

import (
	"fmt"
	"time"
)

func main(){
	timeNow := time.Now()
	fmt.Println(timeNow.Format("20060102"))
}
