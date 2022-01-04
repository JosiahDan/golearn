package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "root:12345@tcp(127.0.0.1:3306)/learning?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	//与数据库建立连接验证信息是否正确
	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Println(err)
	}
}
