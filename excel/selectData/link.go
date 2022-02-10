package main

import (
	"github.com/gin-gonic/gin"
	"golearning/excel/selectData/models"
)

func main() {
	r := gin.Default()
	r.GET("/selectData", func(c *gin.Context) {
		models.SelectData("./excel/selectData/sheets/0116ga/")
		c.JSON(200, models.GetTarGetSheets("./excel/selectData/sheets/output/0116待核查/", "./excel/selectData/sheets/0116ga/1.25日54人.xlsx"))
	})

	r.Run()
}
