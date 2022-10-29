package main

import "fmt"

/*
切片：切片是对数组一个连续片段的引用，所以切片是一个引用类型，这个片段可以是整个数组，或者是由起始和终止索引标识的一些子集
需要注意的是，终止索引的下标项不包含在切片内，切片提供了一个相关数组的动态窗口，容量是动态的，可以通过内建函数cap()获取
到切片的容量，具体切片语法同Python中的切片，比较简单, go中没有负数下标，切片不能用负数，区别于Python
切片追加元素：底层追加元素的时候会对原数组进行扩容，扩容是创建一个新数组，将老数组中的值copy到新数组，老数组不会改变，
扩容后需要使用新的变量接收扩容后的新数组
*/
func main() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	// 切片
	var slice []int = arr[1:3]
	fmt.Printf("切片的类型: %T, 切片的长度: %d, 切片的容量: %d\n", slice, len(slice), cap(slice))
	fmt.Println("slice: ", slice)
	// 不指定初始索引，标识从头开始, 不指定结束索引，表示到末尾
	var slice1 = arr[:3]
	var slice2 = arr[2:]
	fmt.Println(slice1, "\n", slice2)
	// 切片的定义
	// 方式一为上面的切片，通过数组获取
	// 方式二：通过make函数获取切片: 第一个参数为切片类型，第二个是长度，第三个是容量
	var slice3 = make([]int, 4, 20)
	slice3[0] = 11
	slice3[1] = 22
	slice3[2] = 33
	fmt.Println(slice3)
	// 方式三：直接初始化
	var slice4 = []int{1, 4, 5}
	fmt.Println(slice4)
	// 对切片追加元素: 使用内建函数append
	slice4 = append(slice4, 8, 9)
	fmt.Println(slice4)
	slice4 = append(slice4, slice3...) // 追加一个切片，后面的三个点必须写
	fmt.Println(slice4)
	// 切片的copy
	var slice5 = make([]int, 10)
	var slice6 = []int{1, 3, 5, 7}
	copy(slice5, slice6) // 将slice6中的所有元素拷贝到slice5中
	fmt.Println(slice5)
}
