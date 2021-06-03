package main

import ("fmt"; "math")

func main() {
	// 变量声明，如果变量声明时已经初始化，则go可以自动推断其类型, 如果没有指定初始值，默认为0
	var width, height int

	fmt.Println("width is: ", width, "height is: ", height)
	width = 100
	height = 50

	fmt.Println("width is: ", width, "height is: ", height)

	// 同一个语句中声明不同的变量类型, 可以使用换行符或者分号来分隔不同的语句
	var (name = "eason"
		 age = 10
		 address string)

	var (c = 10; d = 20)
	fmt.Println("c is: ", c, "d is: ", d)

	fmt.Println("name is: ", name, "age is: ", age, "address is: ", address)
	address = "陕西西安"
	fmt.Println("name is: ", name, "age is: ", age, "address is: ", address)

	// 简短声明变量, 省略var关键字，赋值使用:=来替代=, 多个变量之间用逗号分隔
	// 使用简短声明时，:=左边的变量都需要有初始值，否则会报错
	a, b := 1, "name"
	fmt.Println("a is: ", a, "b is: ", b)

	// 简短声明的语法要求:=操作符左边至少有一个变量是尚未声明的，否则会报错
	b, e := "20", 20
	fmt.Println("b is: ", b, "e is: ", e)

	// 变量也可以在运行时赋值
	i, h := 23.3, 22.2
	f := math.Max(i, h)
	fmt.Println("f is: ", f)
}
