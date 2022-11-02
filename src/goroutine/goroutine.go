package main

import (
	"fmt"
	"time"
)

/*
go语言中使用go关键字开启一个协程来执行相应的任务
go funcName()
 */
func main() {
	// 开启10个协程计算
	for i := 0; i < 10; i++ {
		go calculate(i, i + 1)
	}
	// 开启10个协程，给切片中添加元素
	slice := make([]int, 4) // 指定切片的长度小于协程的数量
	for i := 0; i < 10; i++ {
		// 当i大于3时，这里会发生异常，需要进行异常处理，否则整个程序都会panic
		go addElement(slice, i)
	}
	// 等待所有的协程执行完毕, 这里需要同步机制，第一种同步方式就是使用time.sleep使主程序进入休眠
	time.Sleep(2 * time.Second)

}

// 协程中执行的函数需要进行异常处理, 使用defer-recover机制
func calculate(m int, n int) {
	defer func() {
		err := recover()
		if nil != err {
			fmt.Println("error: ", err)
		}
	}()
	fmt.Printf("%d + %d = %d\n", m, n, m + n)
}

// 定义一个函数，往切片中添加元素
func addElement(slice []int, num int) {
	// 防止整个程序进入panic，这里需要捕获错误
	defer func() {
		err := recover() // 发生异常之后会将异常信息传递到recover函数
		if nil != err {
			// runtime error: index out of range [n] with length 4
			fmt.Println("add element error: ", err)
		}
	}()
	slice[num] = num // 这里可能会发生异常
	fmt.Println(slice)
}