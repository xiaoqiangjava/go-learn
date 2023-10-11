package main

import (
	"fmt"
	"reflect"
)

/*
go语言中的反射：
go中两个非常重要的函数和类型分别是：
1. reflect.TypeOf() 能获取类型信息，该函数返回的类型是 reflect.Type
2. reflect.ValueOf() 能获取数据的运行时表示，该函数返回的类型是 reflect.Value
 */
func main() {
	author := "Eason"
	// reflect.TypeOf 可以获取变量的类型
	fmt.Println("Type of author: ", reflect.TypeOf(author))
	// reflect.ValueOf 可以获取变量对应的值
	fmt.Println("Value of author: ", reflect.ValueOf(author))
	teacher := Teacher{Name: "Eason", Age: 18}
	rtype := reflect.TypeOf(teacher)
	// 获取类型实现的方法数量
	fmt.Println("Teacher method num: ", rtype.NumMethod())
	method := rtype.Method(0)
	fmt.Println(method)

}

type Teacher struct {
	Name string
	Age int
}

func (teacher Teacher) teacher() {

}

func (teacher *Teacher) eat() {

}

