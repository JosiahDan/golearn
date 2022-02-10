package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
)

//判断数据类型
func getDataType(i interface{}) {
	fmt.Println(reflect.TypeOf(i))
}

func main() {
	f, err := excelize.OpenFile("D:/golearning/excel/readSheet/sheets/213广天所1.7.xlsx")
	if err != nil {
		fmt.Printf("读取表格失败:%s\n", err)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("关闭表格错误:%s\n,", err)
		}
	}()

	rows, err := f.GetRows("待核查人员")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows[2:] {
		getDataType(row)
		fmt.Println(row)
	}

}
