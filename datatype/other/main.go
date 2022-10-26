package main

import (
	"fmt"
	_ "strconv" // _在go中有特殊的含义，表示忽略，包括忽略返回值，包等
)

/*
复杂数据类型：
	指针：里面涉及两个符号&和*，&获取变量对应的指针，即当前变量的内存地址；*获取指针指向的变量的值。指针的本质就是一个内存地址
*/
func main() {
	// 指针数据类型
	var age int = 18
	var ptr *int = &age                                                              // int类型的指针对应的数据类型用*int表示
	fmt.Printf("ptr对应的数据类型为: %T, ptr代表的值为: %v, prt指针对应的变量的值: %v \n", ptr, ptr, *ptr) // ptr对应的数据类型为: *int, ptr代表的值为: 0xc000014098, prt指针对应的变量的值: 18
	fmt.Println("----------------------------")

	// 指针使用注意事项
	// 1. 可以通过指针改变指向的值
	*ptr = 20 // *ptr获取的是age变量对应的值，这里重新赋值，就会修改age变量对应的值，因为age变量的内存地址没有变化，这里只是修改了内存地址对应的值
	fmt.Printf("修改之后age变量对应的值: %v, age变量对应的内存地址为: %v \n", age, &age)
	// 2. 指针变量只能接受一个地址值
	// var ptr1 *int = age 这行代码编译不通过
	// 3. 指针变量的类型必须匹配
	// var ptr1 *float32 = &age
	// 4. 基本数据类型都有对应的指针类型，形式为*基本数据类型，比如：*int, *float etc.
	fmt.Println("----------------------------")

	// go语言标识符，需要注意的就是_在go中有特殊的含义，表示忽略
	/*
		起名规则：
			1. 包名：尽量保持package的名字和目录保持一致，尽量采取有意义的包名，简短、有意义，不要和标准库冲突。
			package main是一个程序的入口，所以你的main函数所在的包就定义为main包，如果不定义为main包，那么就不能得到一个可执行文件
			2. 变量名、函数名、常量名：采用驼峰法
			3. 如果变量名、函数名、常量名首字母大写，则可以被其他包访问，类似Java中的public
	*/

}
