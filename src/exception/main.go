package main

import (
	"errors"
	"fmt"
)

// 使用defer + recover实现异常的捕获
func main() {
	test1()
	fmt.Println("end")
}

func test() {
	defer func() {
		err := recover() // 调用内置函数recover可以捕获程序中遇到的异常
		if nil != err {
			fmt.Println("err: ", err)
		}
	}()
	x := 10
	y := 0
	fmt.Println("x/y=", x/y)
}

// 自定义错误
func test1() (err error) {
	x := 10
	y := 0
	if y == 0 {
		err = errors.New("除数不能为0") // 自定义错误
		panic(err)                 // 内置函数panic会终止程序继续执行，并输出错误信息
	}
	fmt.Println("x/y=", x/y)
	return nil
}
