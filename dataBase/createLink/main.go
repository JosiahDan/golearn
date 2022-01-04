package main

import (
	"database/sql"
	"fmt"
	//引入数据库驱动
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//定义账户名称密码和数据库地址以及数据库名称
	dsn := "root:Djx757198@tcp(127.0.0.1:3306)/learning"
	//返回一个err,如果出错打印出错误
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	//此处defer放在最后是因为如果连接错误程序会直接panic结束程序而不执行defer后的语句
	//defer的语句是在return之前执行
	defer db.Close()
}
