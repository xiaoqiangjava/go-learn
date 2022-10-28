package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串常用函数
func main() {
	var str string = "Golang 你好"
	// 1. len(): 统计字符串的长度，按照字节进行统计，这里要区别与其他语言
	fmt.Printf("str的长度为: %d \n", len(str)) // 13 字节的长度
	fmt.Println("----------------------------")
	// 2. 字符串遍历
	// 方式一：for-range
	for i, val := range str {
		fmt.Printf("下标: %d, 值: %c \n", i, val)
	}
	// 方式二：切片 r := []rune(str)
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Println(r[i]) // 这种方式答应的是字符对应的编码值
		fmt.Printf("%c\n", r[i])
	}
	fmt.Println("----------------------------")
	// 3. 字符串转整数
	// 方式一: strconv.ParseInt() // 这种方式需要三个参数, 转换结果是什么类型有第三个参数决定
	i, err := strconv.ParseInt("55", 10, 64)
	if nil != err {
		fmt.Println("error: ", err)
	} else {
		fmt.Printf("ParseInt转换结果: %d, ParseInt转换结果类型: %T \n", i, i)
	}
	// 方式二: strconv.Atoi() // 只需要一个参数, 返回值是一个int类型，具体位数跟操作系统有关
	i2, err2 := strconv.Atoi("55")
	if nil != err2 {
		fmt.Println("error2: ", err2)
	} else {
		fmt.Printf("Atoi转换结果: %d, Atoi转换结果类型: %T \n", i2, i2)
	}
	fmt.Println("----------------------------")
	// 4. 整数转字符串
	// 方式一: strconv.FormatInt，需要通过参数指定进制
	s := strconv.FormatInt(55, 10)
	fmt.Println("FormatInt转换结果: ", s)
	// 方式二: strconv.Itoa，默认使用10进制
	s1 := strconv.Itoa(55)
	fmt.Println("Itoa转换结果: ", s1)
	fmt.Println("----------------------------")
	// 5. 统计一个字符串中指定子串的个数，这个跟Python中的count函数一样
	count := strings.Count(str, "Go")
	fmt.Println("go出现次数: ", count)
	fmt.Println("----------------------------")
	// 6. 不区分大小写比较字符串 EqualFold
	fmt.Println("go Go不区分大小写比较: ", strings.EqualFold("go", "Go"))
	fmt.Println("----------------------------")
	// 7. 区分大小写比较字符串 ==
	fmt.Println("go Go区分大小写比较: ", "go" == "Go")
	fmt.Println("----------------------------")
	// 8. 查找指定子串在字符串中第一次出现的位置，不存在返回-1，跟Python中的index函数一致
	fmt.Println("Go在str中出现的位置: ", strings.Index(str, "go"))
	fmt.Println("----------------------------")
	// 9. 字符串的替换 Replace
	s2 := strings.Replace(str, "Go", "go", -1) // 后面的2是替换的个数，有多个满足是，替换指定的个数, 全部替换传-1
	fmt.Println("Replace替换后的字符串: ", s2)
	fmt.Println("----------------------------")
	// 10. 字符串的切割
	arr := strings.Split("java-python-go", "-")
	fmt.Println("arr: ", arr)
	fmt.Println("----------------------------")
	// 11. 字符串大小写转换
	fmt.Println("GO转小写: ", strings.ToLower("GO"), "go转大写: ", strings.ToUpper("go"))
	fmt.Println("----------------------------")
	// 12. 去掉字符串两边的空格或者指定字符
	fmt.Println("去掉空格: ", strings.TrimSpace("   go   lang    "))
	fmt.Println("去掉指定字符: ", strings.Trim("~golang~", "~"))
	fmt.Println("去掉左边指定字符: ", strings.TrimLeft("~golang~", "~"))
	fmt.Println("去掉右边指定字符: ", strings.TrimRight("~golang~", "~"))
	fmt.Println("----------------------------")
	// 13. 判断字符串是否以指定的子串开头或者结尾
	fmt.Println("判断开头: ", strings.HasPrefix("Golang", "Go"))
	fmt.Println("判断结尾: ", strings.HasSuffix("Golang", "lang"))
	fmt.Println("----------------------------")
}
