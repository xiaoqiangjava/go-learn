package main

import "fmt"

/**
 	函数的声明：
	func functionName(parameterName type) returnType {
		// 函数体，具体需要实现的功能
	}
	函数的参数和返回值并不是必须的，像main方法就没有参数和返回值, 如果声明时没有指定返回值，那么函数不能有return语句

	注意：
		:=操作符是简短声明语法，不需要指定变量的类型类声明一个变量，需要给变量赋值时，如果变量已经声明，可以直接使用=操作符

	闭包：闭包的本质是一个匿名函数，只是这个匿名函数引用了外部的变量或者参数
*/
func main() {
	// 普通函数
	fmt.Println(calculateBill(2.3, 2))

	// 参数类型相同的函数
	fmt.Println(calculateBill1(2.3, 2)) // 这里2是常量，因此可以赋值给float类型

	// 多个返回值函数, 接收时使用两个变量接收, 如果使用一个变量会报错
	area, perimeter := rectProps(2.1, 4.1)
	fmt.Println("area is: ", area, "perimeter is: ", perimeter)

	// 命名返回值函数，一旦命名了返回值，可以认为这些值在函数第一行进行了变量的声明，后续直接赋值即可，不需要声明
	area, perimeter = rectProps1(2.2, 4.1)
	fmt.Println("are is: ", area, "perimeter is: ", perimeter)

	// 使用空白符，类似于Scala中的占位符来接受不需要使用的参数
	area, _ = rectProps(2.1, 4) // 返回值perimeter被丢弃
	fmt.Println("area is: ", area)

	// 闭包调用: 通过下面的两种调用方式可以看到，想要达到闭包中类似于全局变量的外部引用，需要使用一个变量来接收
	// 闭包返回的匿名函数，否则每次立即执行该函数，相当于每次都声明了外部引用
	f := getSum()     // 这一行代码相当于声明并初始化了getSum中的变量sum，后面的调用都会修改sum的值
	fmt.Println(f(1)) // 1 这里每次调用都会保留f函数外部应用的sum变量值，相当于一个全局变量
	fmt.Println(f(2)) // 3
	fmt.Println(f(3)) // 6
	// 如果直接调用闭包，那么每次都会重置匿名函数外部引用的值
	fmt.Println(getSum()(1)) // 1
	fmt.Println(getSum()(2)) // 2
	fmt.Println(getSum()(3)) // 3

}

/**
计算商品的价格
*/
func calculateBill(price float64, number int) float64 {
	return price * float64(number)
}

/**
如果声明参数时，连续多个参数具有相同的数据类型，那么无须一一列出，只需要在最后一个参数后面指定类型即可
*/
func calculateBill1(price, number float64) float64 {
	return price * number
}

/**
多返回值函数：传入矩形的长和宽，返回面积和周长
*/
func rectProps(length, width float64) (float64, float64) {
	area, perimeter := length*width, (length+width)*2
	return area, perimeter
}

/**
命名返回值函数：一旦命名了返回值，就相当于在函数的第一行进行了变量的声明，后续直接使用即可
命名返回值函数这里需要注意一下，是第一行声明了两个变量，而没有赋值，后面使用这两个变量即可，不需要再次声明
*/
func rectProps1(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return // 会直接将命名返回值变量返回
}

/*
 闭包：是一个匿名函数和匿名函数外部引用构成的一个整体，即这个匿名函数内部使用了外部的变量或者参数
 闭包的返回值就是一个匿名函数，因为在go语言中，函数也是一种数据类型，类似于Scala
*/
func getSum() func(int) int {
	// 这个变量定义在匿名函数外部
	var sum int = 0
	return func(num int) int {
		sum += num
		return sum
	}
}
