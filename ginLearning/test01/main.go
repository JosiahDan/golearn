package main

import (
	"github.com/gin-gonic/gin"
	"golearning/ginLearning/test01/models"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, models.ResultString())
	})
	r.Run()
}
