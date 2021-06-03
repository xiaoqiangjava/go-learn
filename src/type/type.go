package main

import (
	"fmt"
	"unsafe"
)

/**
 * go 语言支持的基本类型
 * 1. bool 表示一个布尔值，值为true或者false
 * 2. 数字类型
 * 		2.1 有符号整型
 *			2.1.1 int8 表示8位有符号整型，大小8位，范围-128~127
 *			2.1.2 int16 表示16位有符号整型，大小16位，范围-32768~32767
 * 			2.1.3 int32 表示32位有符号整型，大小32位，范围-2147483648～2147483647
 * 			2.1.4 int64 表示64位有符号整型，大小64位，范围-9223372036854775808～9223372036854775807
 * 			2.1.5 int 根据不同的底层平台，表示32或者64位整型，除非对整型大小有特定的需求，否则通常应该使用int表示整型，大小在32位操作系统下是32位，在64位操作系统下64位
 * 		2.2 无符号整型
			2.2.1 unit8 表示8位无符号整型，大小8位，范围0~255
			2.2.2 unit16 表示16位无符号整型，大小16位，范围0~65535
			2.2.3 unit32 表示32位无符号整型，大小32位，范围0～4294967295
			2.2.4 unit64 表示64位无符号整型，大小64位，范围0～18446744073709551615
			2.2.5 unit 根据不同的底层平台，表示32位或者64位无符号整型
		2.3 浮点型
			2.3.1 32位浮点型
			2.3.2 64位浮点型
		2.4 复数类型
			2.4.1 complex64: 实部和虚部都是float32类型的复数
			2.4.2 complex128: 实部和虚部都是float64类型的复数
		2.5 byte 是int8的别名
		2.6 rune 是int32的别名
	3. 字符串类型 字符串是字节的集合
	4. 类型转换 go语言中有着非常严格的强类型特征，没有自动类型转换或者自动类型提升
 */
func main() {
	// 1. bool类型
	a := true
	b := false
	fmt.Println("a: ", a, "b: ", b)
	// ||操作符当a或者b有一个为true时，返回true
	d := a || b
	fmt.Println("d is: ", d)
	// &&操作符当a和b都为true时，返回true
	e := a && b
	fmt.Println("e is: ", e)

	// 2. int类型，有符号整型
	var h int = 65
	i := 88
	fmt.Println("h is: ", h, "i is: ", i)
	// 通过unsafe包来对应的类型占用的字节数，1字节=8位, Printf中%T用于打印变量的类型，在64位操作系统下，int占用64位，8个字节
	fmt.Printf("type of i is: %T, size of i is: %d", i, unsafe.Sizeof(i))

	// 3. uint类型，无符号整型
	var aa uint = 32
	fmt.Println("aa is: ", aa)

	// 4. float浮点型, float64是浮点型的默认类型
	ai, bi := 3.4, 3.5
	fmt.Printf("type of ai is: %T, bi: %T", ai, bi)
	fmt.Println()
	fmt.Println("diff is: ", bi - ai)

	// 5. 复数类型
	c1 := complex(8, 9)
	c2 := 8 + 23i
	fmt.Println("sum: ", c1 + c2)

	// 6. string字符串
	first := "zhang"
	last := "eason"
	name := first + " " + last
	fmt.Println("name is: ", name)

	// 7. 类型转换，没有自动类型转换，需要操作时需要把不同的类型转换成相同类型
	cast1 := 2
	cast2 := 2.3
	sum1 := float64(cast1) + cast2
	fmt.Println("sum1 is: ", sum1)

}
