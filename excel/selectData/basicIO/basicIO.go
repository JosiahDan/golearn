package basicIO

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
)

type Person struct {
	Name    string
	Phone   string
	Id      string
	Address string
}

//读取已分拣表格
func ReadSelectedExcle(readPath string) [][]string {
	f, err := excelize.OpenFile(readPath)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return rows
}

//读取未分拣表格表格
func ReadExcel(path string) ([][]string, int) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
	}

	rows, err := f.GetRows("待核查人员")

	index := 0
	for _, row := range rows[1] {
		if string(row) == "联系地址" {
			break
		}
		index++
	}

	return rows, index
}

//写入表格
func WriteExcel(rows [][]string, writePath string, area string) bool {
	if CheckFolder(writePath) != true {
		MkdirFolder(writePath)
	}

	f := excelize.NewFile()
	streamWriter, err := f.NewStreamWriter("Sheet1")
	if err != nil {
		fmt.Println(err)
		return false
	}

	rowID := 0
	for _, row := range rows {
		colID := 0
		rowList := make([]interface{}, len(row))
		for _, value := range row {
			rowList[colID] = value
			colID++
		}
		cell, _ := excelize.CoordinatesToCellName(1, rowID)
		if err := streamWriter.SetRow(cell, rowList); err != nil {
			fmt.Println(err)
		}
		rowID++
	}

	if err := streamWriter.Flush(); err != nil {
		fmt.Println(err)
		return false
	}

	if err := f.SaveAs(writePath + "/" + area + "待核查人员.xlsx"); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

//通过区域分派表格
func SelectArea(rows [][]string, addressIndex int, writePath string) {

	areaList := []string{"安良", "白庙", "茨芭", "东城", "广天", "黄道", "李口", "龙山", "堂街", "王集", "薛店", "姚庄", "渣园", "长桥", "冢头"}
	for _, area := range areaList {
		var resultList [][]string
		for _, row := range rows[1:] {
			match, _ := regexp.MatchString(area, row[addressIndex])
			if match {
				resultList = append(resultList, row)
				fmt.Println(row)
			}
		}
		WriteExcel(resultList, writePath+"0116待核查", area)
	}
}

//获取目标文件夹下的所有文件名
func GetFileNames(readPath string) []string {
	//创建字符型数组存储文件名
	var resultList []string
	//使用ioutil包获取文件夹下文件的地址并将其文件名添加进resultList
	fileNames, _ := ioutil.ReadDir(readPath)
	fmt.Println(reflect.TypeOf(fileNames))
	for _, i := range fileNames {
		resultList = append(resultList, i.Name())
	}

	return resultList
}

//检查文件夹是否存在,返回值为bool
func CheckFolder(path string) bool {
	//判断文件夹是否存在
	_, err := os.Stat(path)
	if err != nil {
		//不存在则抛出错误
		fmt.Println(err)
		return false
	}
	return true
}

//创建文件夹本地,返回值为bool
func MkdirFolder(path string) bool {
	//创建文件夹、ModePerm为操作权限
	err := os.Mkdir(path, os.ModePerm)

	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println(path + "文件夹创建成功")
	return true
}

func GetDataIndex(path string, key string) int {
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	rows, err := f.GetRows("待核查人员")
	index := 0
	for _, value := range rows[1] {
		if match, _ := regexp.MatchString(key, value); match {
			return index
		}
		index++
	}
	return 0
}

//获取目标文件夹下所有excel表格的数据,并返回JSON数据
func GetSheetDatas(targetPath string, readPath string) []Person {
	var resultList []Person
	nameIndex := GetDataIndex(readPath, "姓名")
	phoneIndex := GetDataIndex(readPath, "手机号")
	idIndex := GetDataIndex(readPath, "证件号")
	addressIndex := GetDataIndex(readPath, "地址")
	if !CheckFolder(targetPath) {
		fmt.Println("文件夹不存在")
		return nil
	}

	fileNames := GetFileNames(targetPath)

	for _, file := range fileNames {
		rows := ReadSelectedExcle(targetPath + file)
		for _, row := range rows {
			var result Person
			result.Name = row[nameIndex]
			result.Phone = row[phoneIndex]
			result.Id = row[idIndex]
			result.Address = row[addressIndex]

			resultList = append(resultList, result)
			fmt.Println("插入中")
		}
	}

	return resultList
}
