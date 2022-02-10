package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func writrDataFile(fileName string) {
	writeFile, err := os.OpenFile("./goroutine/test4/writeFiles/"+fileName+".txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer writeFile.Close()
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().Unix())
		writer := bufio.NewWriter(writeFile)
		_, err := writer.WriteString(fmt.Sprintf("%d\n", rand.Intn(1000)))
		if err != nil {
			fmt.Println("转换错误")
			return
		}
		writer.Flush()
		fmt.Println("写入成功")
	}
	wg.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writrDataFile(fmt.Sprintf("%d", i))
	}
	wg.Wait()
}
