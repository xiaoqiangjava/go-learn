package main

import "fmt"

func main() {
	// 获取用户终端输入, 需要注意的是传入的必须是接受变量的内存地址，否则不会获取到值
	// 方式一：Scanln
	var age int
	fmt.Println("请输入年龄: ")
	fmt.Scanln(&age)

	var name string
	fmt.Println("请输入姓名: ")
	fmt.Scanln(&name)

	var score float64
	fmt.Println("请输入成绩: ")
	fmt.Scanln(&score)

	var isVIP bool
	fmt.Println("请输入是否VIP: ")
	fmt.Scanln(&isVIP)

	fmt.Printf("年龄: %d, 姓名: %s, 成绩: %.2f, 是否VIP: %t \n", age, name, score, isVIP)

	// 方式二：Scanf
	fmt.Println("请输入年龄, 姓名, 成绩, 是否VIP, 用空格分隔: ")
	fmt.Scanf("%d %s %f %t", &age, &name, &score, &isVIP)
	fmt.Printf("年龄: %d, 姓名: %s, 成绩: %.2f, 是否VIP: %t \n", age, name, score, isVIP)
}
