package main

import (
	"fmt"
	"reflect"
)

type myInt int64

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type is %v  kind is %v\n", t.Name(), t.Kind())
}

func main() {
	//指针类型
	var a *float32
	//自定义类型
	var b myInt
	//类型别名
	var c rune
	reflectType(a)
	reflectType(b)
	reflectType(c)

	type person struct {
		name string
		age  int
	}

	type book struct {
		title string
	}

	var d = person{
		name: "张三",
		age:  18,
	}

	var e = book{title: "新华字典"}

	reflectType(d)
	reflectType(e)
}
