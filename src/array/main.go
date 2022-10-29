package main

import "fmt"

/*
数组的定义：var name [n]type，其中name是变量名称，n是数组的大小，type为数组里面存储的数据类型
数组的主意事项：
1. 数组的长度是数组类型的一部分 [4]int
2. go语言中数组是值类型，默认使用的是值传递, 这里区别与大部分的语言，编写代码时需要额外注意，想要修改数组的值是，可以使用指针
*/
func main() {
	var scores [5]int
	scores[0] = 90
	scores[1] = 88
	scores[2] = 79
	scores[3] = 38
	scores[4] = 68
	// 遍历数组
	var sum int
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	fmt.Println("sum: ", sum)
	// 手动输入数组的值
	var scoreArr [5]int16
	for i := 0; i < len(scoreArr); i++ {
		fmt.Printf("请输入第%d个学生的成绩: ", i+1)
		fmt.Scanln(&scoreArr[i])
	}
	// 使用for-range遍历数组
	for i, val := range scoreArr {
		fmt.Printf("第%d个学生的成绩为%v\n", i+1, val)
	}
	// 数组的初始化
	// 方式一：
	var arr1 [3]int = [3]int{3, 6, 9}
	fmt.Println(arr1)
	fmt.Printf("数组的类型: %T\n", arr1)
	// 方式二：
	var arr2 = [3]int{1, 4, 7}
	fmt.Println(arr2)
	fmt.Printf("数组的类型: %T\n", arr2)
	// 方式三：
	var arr3 = [...]int{22, 33, 44, 55, 66} // ...表示数组的长度由后面元素的个数确定
	fmt.Println(arr3)
	fmt.Printf("数组的类型: %T\n", arr3)
	// 方式四：
	var arr4 = [...]int{1: 33, 0: 22, 2: 44}
	fmt.Println(arr4)
	fmt.Printf("数组的类型: %T\n", arr4)

	// go 语言中数组是值类型
	test(arr1)        // 在函数test中对数组做了修改
	fmt.Println(arr1) // 这里输出的还是原来的数组
	// 如果需要修改arr的值，这里的参数需要传递的是arr的指针
	test1(&arr1)
	fmt.Println(arr1)
	fmt.Println("-------------------")

	// 二维数组
	var arr5 [2][3]int
	fmt.Println(arr5)
	var arr6 [2][2]int = [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr6)
	// 二维数组的遍历，使用双层for循环
	for i, v := range arr6 {
		for j, val := range v {
			fmt.Printf("arr[%v][%v]=%v \t", i, j, val)
		}
		fmt.Println()
	}
}

func test(arr [3]int) {
	arr[1] = 10
}

func test1(arr *[3]int) {
	arr[1] = 10 // 访问指针中的值使用*arr，但是这里不适用*也可以修改数组中对应的值
	(*arr)[2] = 20
}
