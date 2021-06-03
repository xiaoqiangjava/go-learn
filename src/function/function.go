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
	area, perimeter := length * width, (length + width) * 2
	return area, perimeter
}

/**
	命名返回值函数：一旦命名了返回值，就相当于在函数的第一行进行了变量的声明，后续直接使用即可
 */
func rectProps1(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return // 会直接将命名返回值变量返回
}