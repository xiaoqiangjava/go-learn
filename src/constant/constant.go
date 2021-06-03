package main

import (
	"fmt"
)

func main() {
	var a string = "I love go"
	fmt.Println("a is: ", a)
	a = "I love python"
	fmt.Println("a is: ", a)
	// 常量使用关键字const声明，声明之后不能再次赋值
	const hello = "Hello World"
	fmt.Println("hello is: ", hello)
	// 常量的值在编译是确定，因为函数调用发生在运行时，所以不能将函数的返回值赋值给常量
	//var b = math.Sqrt(4)
	//const c = math.Sqrt(8)

	// go语言中不允许将一种类型的变量赋值给另一种类型的变量, 但常量时没有指定类型的，所以可以赋值给其他类型的变量
	var defaultName = "Eason"
	type str string // type关键字创建一个新的类型，底层的类型是string
	var customName str = "Eason"  // "Eason"常量是无类型的，可以赋值给任何字符串变量
	const constName = "Eason"
	customName = constName // 常量可以赋值给其他类型的变量
	// customName = defaultName // 报错，类型不同，不能赋值，defaultName现在的类型是string
	fmt.Println("default name is: ", defaultName, "custom name is: ", customName)
}
