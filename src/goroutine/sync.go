package main

import (
	"fmt"
	"sync"
)

/*
go协程之间的同步机制：开启协程之后，相关任务会异步执行，如果我们想要等待所有的任务执行完成，go提供了sync包和channel来解决同步问题，
当然如果你能精准的预测到每个任务的执行时间，也可以通过time.sleep来实现同步
第一种方式：sync包下面的WaitGroup同步goroutine
WaitGroup等待一组goroutine执行完毕，主程序调用Add方法增加等待goroutine的数量，每个goroutine执行结束后调用Done方法，此时等待的
goroutine数量-1，主程序通过Wait阻塞，直到等待队列中goroutine的数量为0，这个跟Java中的CountDownLatch一致
第二种方式：channel实现goroutine之间的同步
通过channel能够在多个goroutine之间通讯，当一个goroutine完成之后想channel发出退出信号，等所有的goroutine退出的时候，利用
for循环去channel中读取结束信号，利用取不到数据会阻塞的原理，等待所有的goroutine执行完毕，使用这种方式的前提条件是你已经知道
自己开启了多少个goroutine
 */
func main() {
	// 方式一：开启协程，使用sync.WaitGroup控制goroutine之间的同步
	var wait sync.WaitGroup // 声明一个sync.WaitGroup类型的变量
	for i := 0; i < 10; i++ {
		wait.Add(1) // WaitGroup队列中增加goroutine
		go cal(i, i + 1, &wait)
	}
	wait.Wait() // 等待所有的goroutine执行完毕
	fmt.Println("[sync] all goroutines done")

	// 方式二：通过channel实现goroutine之间的同步
	var channel = make(chan bool, 10)
	for i := 10; i < 20; i++ {
		go cal1(i, i + 1, channel)
	}
	// 遍历channel，获取数据，利用从channel中读取数组会阻塞的原理等待所有goroutine执行完毕
	for i := 0; i < 10; i++ {
		<-channel // 读取到的数据不使用可以直接丢弃
	}
	close(channel) // 关闭管道
	fmt.Println("[channel] all goroutines done")
}

func cal(m int, n int, wait *sync.WaitGroup) {
	defer func() {
		err := recover()
		if nil != err {
			fmt.Println("calculate error")
		}
	}()
	fmt.Printf("%d + %d = %d\n", m, n, m + n)
	wait.Done() // wait.Done()方法会使WaitGroup队列中goroutine的数量-1
}

func cal1(m int, n int, channel chan bool) {
	defer func() {
		err := recover()
		if nil != err {
			fmt.Println("calculate error!")
		}
	}()
	fmt.Printf("%d + %d = %d\n", m, n, m + n)
	channel <- true
}

