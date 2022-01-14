package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("./sheets/213广天所1.7.xlsx")
	if err != nil {
		fmt.Printf("读取表格失败:%s\n", err)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("关闭表格错误:%s\n,", err)
		}
	}()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

}
