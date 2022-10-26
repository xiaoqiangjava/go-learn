package main

import "fmt"

// 全局变量的声明
var x1, x2 = 23, "Eason"

var (
	xx = 11
	yy = 22
)

func main() {
	// 1. 变量的声明
	var age int
	// 2. 变量的赋值
	age = 30
	// 3. 变量的使用
	fmt.Println("age: ", age)
	// 4. 变量声明的同时赋值
	var name string = "Eason"
	fmt.Println("name: ", name)

	// 变量的使用方式
	// 1. 第一种使用方式
	var num int = 18
	fmt.Println("num: ", num)

	// 2. 第二种使用方式, 指定变量的类型，但是不赋值，使用默认值，int类型的默认值是0
	var n int
	fmt.Println("n: ", n)

	// 3. 第三种使用方式，不指定具体的变量类型时，会进行自动类型推断
	var m = 18
	fmt.Println("m: ", m)

	// 4. 第四种使用方式，可以省略var关键字，使用:=给变量赋值
	x := 20
	fmt.Println("x: ", x)

	var (
		y1 = 11
		y2 = "Eason"
	)
	fmt.Println("y1: ", y1, "y2: ", y2)

	// 多变量声明
	var n1, n2, n3 int
	fmt.Println("n1: ", n1, "n2: ", n2, "n3: ", n3)
	var sName, sAge, sScore = "Eason", 29, 12.12
	fmt.Printf("sName: %s, sAge: %d, sScore: %.2f", sName, sAge, sScore)
	fmt.Println()

	// 使用全局变量
	fmt.Printf("x1: %d, x2: %s", x1, x2)
	fmt.Println()
	fmt.Println("xx: ", xx, "yy: ", yy)

}
