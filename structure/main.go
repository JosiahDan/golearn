package main

import (
	"encoding/json"
	"fmt"
)

//学生结构体
type student struct {
	ID    string
	Name  string
	Age   int8
	Score int8
}

//学生构造函数
func createStudent(id string, name string, age int8, score int8) *student {
	return &student{
		ID:    id,
		Name:  name,
		Age:   age,
		Score: score,
	}
}

//班级结构体
type class struct {
	Title    string
	Students []*student
}

//class方法展示全部学生信息
func (c class) showAllInformation() {
	if c.Students == nil {
		fmt.Println("该班级信息为空")
		return
	}
	for _, value := range c.Students {
		data, err := json.Marshal(value)
		if err != nil {
			fmt.Println("json序列化失败")
		}
		fmt.Printf("%s\n", data)
	}
}

//class方法添加学生信息
func (c *class) addStudent() {
	var id, name string
	var age, score int8
	fmt.Println("输入学生ID")
	fmt.Scanln(&id)
	for _, value := range c.Students {
		if value.ID == id {
			fmt.Println("该学生已存在")
			return
		}
	}

	fmt.Println("输入学生姓名")
	fmt.Scanln(&name)
	fmt.Println("输入学生年龄")
	fmt.Scanln(&age)
	fmt.Println("输入学生成绩")
	fmt.Scanln(&score)

	c.Students = append(c.Students, createStudent(id, name, age, score))
}

//class方法查询单个学生信息
func (c class) getStudent() *student {
	var id string
	fmt.Println("输入要查询的学生ID")
	fmt.Scanln(&id)
	for _, value := range c.Students {
		if value.ID == id {
			data, err := json.Marshal(value)
			if err != nil {
				fmt.Println("json序列化失败")
			}
			fmt.Printf("%s\n", data)
			return value
		}
	}
	fmt.Println("查无此人")
	return nil
}

//class方法修改学生信息
func (c *class) editStudent() {
	newInformation := c.getStudent()
	if newInformation == nil {
		return
	}
	fmt.Println("新姓名")
	fmt.Scanln(&newInformation.Name)
	fmt.Println("新年龄")
	fmt.Scanln(&newInformation.Age)
	fmt.Println("新成绩")
	fmt.Scanln(&newInformation.Score)
	fmt.Println("修改完成")
}

//class方法删除学生信息
func (c *class) deleteStudent() {
	flag := 0
	var id string
	fmt.Println("输入要删除的学生id")
	fmt.Scanln(&id)
	for _, value := range c.Students {
		if value.ID == id {
			c.Students = append(c.Students[:flag], c.Students[flag+1:]...)
			fmt.Println("删除成功")
			return
		}
		flag++
	}
	fmt.Println("查无此人")
}

func menu(class1 class) {
	for {
		fmt.Printf("学生信息管理系统\n")
		fmt.Printf("1.展示班级学生列表\n2.查询学生信息\n3.新增学生\n4.编辑学生\n5.删除学生\n6.退出\n")
		var key int8
		fmt.Scanln(&key)
		switch key {
		case 1:
			class1.showAllInformation()
		case 2:
			class1.getStudent()
		case 3:
			class1.addStudent()
		case 4:
			class1.editStudent()
		case 5:
			class1.deleteStudent()
		case 6:
			return
		}
	}
}

func main() {
	var class01 class
	menu(class01)
}
