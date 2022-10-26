package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

/*
基本数据类型
数值型：
	整数类型：int, int8, int16, int32, int64; 整型的默认值为0
	浮点型：float32, float64; 浮点型的默认值为0
字符型：没有单独的字符类型，使用byte来保存单个的ASCII码字符，使用int或更大的整型表示Unicode编码;
布尔型：bool; 布尔类型的默认值为false
字符串：string 在go语言中，string属于基本数据类型，首字母是小写的，区别于Java; 字符串的默认值为""
*/

func main() {
	// int, int8, int16, int32, int64 整数数据类型
	// fmt格式化时，%T可以打印变量对应的数据类型，在go语言中，默认使用的整数类型是int，在32位操作系统中为4个字节，在64位操作系统占8个字节
	var age = 18
	fmt.Printf("age的数据类型为: %T", age)
	fmt.Println()
	fmt.Println("int类型占用字节: ", unsafe.Sizeof(age))
	fmt.Println("----------------------------")

	// float32, float64是go语言中的浮点类型, 在go语言中，默认的浮点类型为float64
	// 浮点类型会有精度的丢失，float64表示的数精度更高，因此建议使用float64
	var num float32 = 3.14
	fmt.Println("浮点数: ", num)
	var num1 = 3.14
	fmt.Printf("go默认的浮点类型为: %T", num1)
	var x float32 = 33.333333
	var y float64 = 33.333333
	fmt.Println()
	fmt.Println("x: ", x, "y: ", y)                                                       // 输出：x:  33.333332 y：  33.333333   float32表示的浮点数精度有丢失
	fmt.Println("float32占用的字节数: ", unsafe.Sizeof(x), "float64占用的字节数: ", unsafe.Sizeof(y)) // float32占4个字节，float64占8个字节，与底层是32位还是64位操作系统无关，区别与int
	fmt.Println("----------------------------")

	// 字符类型: 字符类型本质上就是一个整数，也可以直接参与运算，输出字符的时候，会将对应的ASCII码值作为一个输出
	// golang中没有具体的字符类型，但是可以使用byte来代替，区别与Java中的char
	var c1 byte = 'a'
	fmt.Println("c1: ", c1)
	var c2 byte = '6'
	fmt.Println("c2: ", c2)
	// 英文符号，标点对应的是ASCII码值，而汉字对应的是Unicode码值，但本质字符还是会转换成数字来输出，因此中文字符可以使用int或者更大类型的整形来存储
	var c3 int = '中'
	fmt.Println("汉字字符c3: ", c3) // 这里输出20013
	// 总结：Golang的字符对应使用的是UTF-8编码，ASCII码是Unicode的前128位
	// 如果要输出对应字符的具体值，可以使用格式化输出来完成
	var c4 = '中'
	fmt.Printf("c4对应的字符值为: %c", c4)
	fmt.Println()
	fmt.Println("----------------------------")

	// 转义字符: \n 换行
	fmt.Println("aaa\nbbb")
	// 转义字符: \b 退格
	fmt.Println("aaa\bbbb") // 输出aabbb, 因为\b是退格，将最后一个a覆盖掉
	// 转义字符: \r 光标回到本行的开头，后续就会替换原有的字符
	fmt.Println("aaaaa\rbbb") // 输出bbbaa, \r光标回到开头，后面输出的bbb会替换掉前面的三个a
	// 转义字符: \t 制表符，每8位构成一个制表符，不够8位时用空格自动补齐
	fmt.Println("aaaa\tbbb")
	// 转义字符: \"  原样输出双引号
	fmt.Println("\"Golang\"") // 输出"Golang"
	fmt.Println("----------------------------")

	// 布尔类型，bool类型，bool类型数据只允许取值true和false
	var flag bool = true
	fmt.Println("flag: ", flag)
	flag = false
	fmt.Println("flag: ", flag)
	fmt.Println("5 < 9: ", 5 < 9)
	fmt.Println("----------------------------")

	// 字符串类型：一串固定长度的字符连接起来的字符序列, 字符串是不可变的
	var s1 string = "Eason"
	fmt.Println("字符串：", s1)
	// 字符串中如果存在特殊字符，使用``引号来定义
	var s2 string = `Eason, hello "Golang"`
	fmt.Println("带特殊字符的字符串: ", s2)
	// 字符串拼接，当一个字符串过长时，由于go语言省略;的原因，换行时需要将+留在上一行
	var s3 = "Eason, hello" + "Eason, hello" + "Eason, hello" + "Eason, hello" +
		"Eason, hello" + "Eason, hello" + "Eason, hello" + "Eason, hello"
	fmt.Println("字符串拼接：", s3)
	fmt.Println("----------------------------")

	// 基本类型的默认值
	var a int
	var b float32
	var c float64
	var d bool
	var e string
	fmt.Println("a: ", a, "b: ", b, "c: ", c, "d: ", d, "e: ", e) // a:  0 b:  0 c:  0 d:  false e:
	fmt.Println("----------------------------")

	// 类型转换：go语言中没有隐式类型转换，发生类型转换时只能强制转换
	var n1 int = 100
	var n2 float32 = float32(n1)
	fmt.Println("类型转换：", n2)
	// 将int64转换为int8时，编译不会报错，但是会出现数据溢出
	var n3 int64 = 888888888888
	var n4 int8 = int8(n3)
	fmt.Println("类型向下转换时会溢出：", n4) // 这里输出56
	// go语言不存在自动类型转换，因此不能将两个不同类型的整数计算的结果直接赋值给不匹配的数据类型, 编译会报错
	var n5 int8 = 23
	var n6 int64 = int64(n5) + 33 // 一定要匹配=左右的数据类型
	fmt.Println("自动类型转换测试：", n6)
	// 强制类型转换在计算得到结果的情况下，如果计算之前已经超出了该类型的范围，编译会报错，如果是计算之后得到的结果会超出范围，编译不会报错，数据会溢出
	// var n7 int8 = int8(12) + 116 这里编译直接报错，应该是在高版本的编译器里面做了优化，计算结果超出范围编译不通过
	fmt.Println("----------------------------")

	// 基本类型转string
	// 方式一：使用fmt包下的Sprintf，第一个参数是格式，不同的数据类型，使用不同的标识表示，具体参考文档
	// 常用的%d表示十进制整数，%t表示bool数, %c表示字符, %f表示浮点数, %v值的默认格式表示
	var xx int = 19
	var str = fmt.Sprintf("%d", xx)
	fmt.Printf("str的数据类型为: %T, str对应的值为: %v \n", str, str)
	var x2 bool = false
	var str2 = fmt.Sprintf("%t", x2)
	fmt.Printf("str2的数据类型为: %T, str2对应的值为: %v \n", str2, str2)
	// 方式二：使用strconv包下面的Format*函数来实现
	var str3 = strconv.FormatInt(int64(xx), 10) // 第一个参数必须转为int64类型，第二个参数是字面值的进制形式
	fmt.Printf("str3的数据类型为: %T, str3对应的值为: %v \n", str3, str3)
	var x4 = 3.14
	var str4 = strconv.FormatFloat(x4, 'f', 2, 64)
	fmt.Printf("str4的数据类型为: %T, str4对应的值为: %v \n", str4, str4)
	fmt.Println("----------------------------")

	// string类型转为基本类型，使用strconv包下面的Parse*函数
	var s5 = "true"
	// func ParseBool(str string) (value bool, err error)  后面的是返回值
	var b1, _ = strconv.ParseBool(s5)
	fmt.Printf("b1的数据类型为: %T, b1对应的值为: %v \n", b1, b1)
	var s6 = "18"
	// func ParseInt(s string, base int, bitSize int) (i int64, err error)
	var i1, _ = strconv.ParseInt(s6, 10, 64)
	fmt.Printf("i1的数据类型为: %T, i1对应的值为: %v \n", i1, i1)
	// 注意：string类型转换为基本类型时，如果对应值无法转换，则转换结果为对应值的默认值
	var s7 = "golang"
	var i2, _ = strconv.ParseBool(s7)
	fmt.Printf("i2的数据类型为: %T, i2对应的值为: %v \n", i2, i2)

}
