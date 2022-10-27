package main

import "fmt"

// 流程控制
func main() {
	var count = 20
	// 1. 单分支，if后面的括号可以不写，但是go语言中建议不写(), 但是go语言中的{}是不能省略的
	if count < 30 {
		fmt.Println("库存不足")

	}
	// golang中，if后面可以丙类的加入变量的定义，使用;分隔, 需要注意的是这里定义变量不能使用var关键字
	if num := 10; num < 20 {
		fmt.Println("库存不足")
	}
	// 2. 双分支
	if count > 20 {
		fmt.Println("库存充足")
	} else { // 这里需要注意的是，go语言中，else是不能写到下一行的，有固定的格式限制
		fmt.Println("库存不足")
	}
	// 3. 多分支
	var score = 88
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else {
		fmt.Println("及格")
	}

	// 4. switch分支, go语言中的case语句中不需要写break语句，只会匹配一个
	var level = "A"
	switch level { // switch后面是一个表达式，可以是变量，有返回值的函数
	case "A":
		fmt.Println("成绩大于90")
		fallthrough // case穿透，相当于走了A之后还会继续往后面执行
	case "B":
		fmt.Println("成绩大于80")
	case "C", "D": // case后面可以匹配多个值，使用逗号分隔
		fmt.Println("成绩大于60")
	default: // default可有可无，而且位置也随意
		fmt.Println("成绩小于60")
	}
	// 5. switch后面可以不使用表达式，这个时候当做if使用
	switch {
	case level == "A":
		fmt.Println("成绩大于90")
	case level == "B":
		fmt.Println("成绩大于80")
	default:
		fmt.Println("成绩小于80")
	}
	// 6. switch后面可以定义变量，但是必须以;结尾，负责不通过,这种写法也是当做if来使用
	switch num := 20; {
	case num == 20:
		fmt.Println("num = 20")
		fallthrough // 这里也可以使用fallthrough关键字，case穿透
	default:
		fmt.Println("num != 20")
	}
	fmt.Println("---------------------------")

	// 7. 循环结构, 也是不需要写()，这里的变量定义不能使用var关键字来定义，需要使用:=
	for i := 0; i < 10; i++ {
		fmt.Println("i = : ", i)
	}
	// 8. 循环的结构非常灵活，与Java不同的是省略了;;
	i := 1
	for i < 5 {
		fmt.Println("Golang")
		i++
	}
	// 9. 死循环
	// for {
	// 	fmt.Println("死循环")
	// }

	// 10. for range
	// 普通for循环遍历字符串
	var str = "Golang" // 使用len时，这里暂时不能使用中文
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i])
	}
	// 使用for range, range会将下标赋值给第一个变量，下标对应的值复制到第二个变量，类似于Java中的map循环
	// for range可以使用中文，中文占3个字节，对应的下标是将其第一个下标赋值给第一个变量
	for i, val := range str {
		fmt.Printf("下标: %d, 对应的值: %c \n", i, val)
	}
}
