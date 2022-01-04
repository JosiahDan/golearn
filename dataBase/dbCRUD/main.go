package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "root:Djx757198@tcp(127.0.0.1:3306)/learning?charset=utf8mb4&parseTime=True"

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

type user struct {
	id   int
	age  int
	name string
}

//查询单行数据
func queryRow() {
	//通过id查询表中特定行的数据
	sqlStr := "select id, name, age from user where id=?"
	var u user
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("id:%d,name:%s,age:%d\n", u.id, u.name, u.age)
}

//多行查询
func queryMultiRow() {
	//通过id查询表中特定行的姓名年龄,注意此处为大于某个id而不是等于某个id故为查询大于某个id的所有行
	sqlStr := "select id, name, age from user where id>?"

	//查询大于0的所有行的数据
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	//最后释放rows的所有数据库连接
	defer rows.Close()

	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("id:%d,name:%s,age:%d\n", u.id, u.name, u.age)
	}

}

//插入数据(插入，更新，删除操作均使用Exec方法)
func insertRow() {
	//插入user表并包含姓名年龄两个数据
	sqlStr := "insert into user(name,age) value (?,?)"
	ret, err := db.Exec(sqlStr, "王麻", 28)
	if err != nil {
		fmt.Println(err)
		return
	}

	//在表的末尾插入新的数据，并返回其id
	theId, err := ret.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("插入成功,其id为%d", theId)
}

//更新数据(更新某一行数据的数值)
func updateRow() {
	//通过id去更新某个行的年龄
	sqlStr := "update user set age =? where id =?"
	//更新第1行的年龄为39
	ret, err := db.Exec(sqlStr, 39, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	//n为更新了几行
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("第%d行更新成功\n", n)
}

//删除某行数据
func deleteRow() {
	//删除user表中的某一行(通过id定位)
	sqlStr := "delete from user where id=?"
	//通过id删除某行数据
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	//返回操作操作了几行
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("第%d行删除成功\n", n)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println(err)
	}
}
