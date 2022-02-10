package models

import (
	"golearning/excel/selectData/basicIO"
)

//分拣乡镇
func SelectData(readPath string) {
	//readPath := "./excel/selectData/sheets/0116ga/"
	writePath := "./excel/selectData/sheets/output/"
	//获取文件下所有文件的名称
	fileNames := basicIO.GetFileNames(readPath)

	for _, file := range fileNames {
		rows, addressIndex := basicIO.ReadExcel(readPath + file)
		basicIO.SelectArea(rows, addressIndex, writePath)
	}
}

//获取目标文件夹下的数据
func GetTarGetSheets(readPath string, oldReadPath string) []basicIO.Person {
	result := basicIO.GetSheetDatas(readPath, oldReadPath)
	return result
}
