package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type user struct {
	ID       string `db:"ID"`
	Password string `db:"password"`
	Level    string `db:"level"`
}

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/jxpandemic?charset=utf8mb4&parseTime=True"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

// 查询单条数据示例
func queryRowDemo() {
	sqlStr := "select ID, password, level from user where ID=?"
	var u user
	err := db.Get(&u, sqlStr, "admin")
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%s name:%s age:%s\n", u.ID, u.Password, u.Level)
}

func main() {
	initDB()
	queryRowDemo()
}
