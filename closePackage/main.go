package main

import (
	"fmt"
	"strings"
)

//闭包:匿名函数加上引用外部变量
//判断闭包的方法:观察匿名函数中是否引用了外部变量
//闭包实例
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		//HasSuffix是strings包中的方法用来判断一个字符型变量的结尾是否是由给定字符结尾
		//例如以下例子中 name变量为待判断字符,suffix为结尾字符,如果name末尾是以suffix结尾则返回true反之返回false
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	//此处jpgFunc被赋值了一个函数
	jpgFunc := makeSuffixFunc(".jpg")
	fmt.Println(jpgFunc("jojo"))
}
