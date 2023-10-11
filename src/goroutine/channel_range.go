package main

import (
	"fmt"
	"time"
)

/*
channel的遍历：
 */
func main() {
	// 定义一个输出channel，只有一个buffer，不关闭的话，for-range循环不会退出
	channel := make(chan int, 1)
	// 给channel中写数据
	go write(channel)
	// 遍历
	for val := range channel {
		fmt.Println("val: ", val)
	}
}

// 往channel中写数据
func write(write chan<- int) {
	defer close(write)
	for {
		write <- 1
		time.Sleep(time.Second * 1)
	}
}