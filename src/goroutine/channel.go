package main

import (
	"fmt"
	"time"
)

/*
go语言中是通过goroutine是实现并发的，使用go关键字让一个函数在goroutine中执行，多个goroutine之间通信使用channel来交换数据
go funcName()  开启一个协程来执行函数funcName
当启动多个goroutine时，如果其中的一个goroutine异常了，并且我们并没有对异常进行处理，那么整个程序都会终止，所以我们在编写程序的
时候最好每个goroutine执行的函数都做异常处理，异常处理使用refer+recover实现
channel有三种：
只读channel：<-chan int
只写channel: chan<- int
读写channel: chan int
但是需要注意如果一个管道只是读取数据，没有写入数据的话，会引发deadlock异常，因此定义了只读管道是没有意义的
close(channel) 方法可以关闭通道，但是需要注意的是这里只能关闭双向通道或者仅发送类型的channel，只读channel不能关闭

channel根据是否带有容量分为两种：带缓冲区的channel和不带缓冲区的channel
make(chan int, 10) // 带缓冲区，如果缓存对列是满的，那么发送操作将阻塞，相反如果队列是空的，接受操作会阻塞，言外之意带缓冲区的channel我们可以在不阻塞的情况下连续向该channel中发送缓冲区大小的值
make(chan int) // 不带缓冲区，不带缓冲区的channel在发送发发送完数据后会阻塞，知道接收到接受了数据，当前goroutine才会继续往下执行,不带缓冲区的channel也叫做同步channel
通常业务场景中都会存在在一个函数中当前channel是只读的，另外一个函数中channel是只写的，对于这种场景，我们只需要在传参时指定当前channel是
用于读还是写即可
注意：
channel中的数据必须是成对出现的，有发送，有读取，如果直接收，没有发送或者只发送没有接收，那么所有的goroutine无法运行，程序将会以deadlock结束
因此在定义channel时，如果定义了只读、只写chan，是没有意义的，程序会以deadlock结束
channel主要用来协程之间通信，因此发送方和接收方一般都是不同的协程
 */
func main() {
	// 定义
	//define()
	// 测试不带缓冲区的channel阻塞
	noBufferChan := make(chan int)
	go noBufferSend(noBufferChan)
	go noBufferReceive(noBufferChan)
	time.Sleep(2 * time.Second)
	fmt.Println("----------------------")

	// 带缓冲区的channel
	bufferChan := make(chan int, 3)
	go bufferSend(bufferChan)
	time.Sleep(2 * time.Second)
	go bufferReceive(bufferChan)
	time.Sleep(4 * time.Second)
	fmt.Println("----------------------")

	// 使用缓冲区大小为1的channel实现锁机制： 同一个协程发送数据，处理逻辑，接受数据
	num := 0
	lock := make(chan int, 1)
	for i := 0; i < 10; i++ {
		go channelLock(&num, lock)
	}
	time.Sleep(2 * time.Second)
}

func channelLock(num *int, lock chan int) {
	// 获取锁，相当于给channel中发送一条数据，因为容量只有1，所以其他协程获取时会阻塞
	lock <- 1
	fmt.Println("逻辑处理")
	*num++
	fmt.Println("num = ", *num)
	// 释放锁
	<- lock
}

// 带缓冲区的channel发送和接收操作在对列没有满时不会阻塞演示
func bufferSend(buffer chan<- int) {
	// 容量未满不阻塞
	fmt.Println("start buffer send: ", time.Now())
	buffer <- 1
	fmt.Println("1 send success")
	buffer <- 2
	fmt.Println("2 send success")
	buffer <- 3
	fmt.Println("3 send success")
	fmt.Println("no used buffer", time.Now())
	// 容量满了之后阻塞
	buffer <- 4
	fmt.Println("end buffer send: ", time.Now())

}

func bufferReceive(buffer <-chan int) {
	// 接受数据
	<- buffer
	<- buffer
	<- buffer
	<- buffer
	fmt.Println("receive end")
}

// 不带缓冲区的channel发送和接收操作会阻塞演示
func noBufferSend(noBuffer chan int) {
	// 发送数据到channel
	fmt.Println("start send: ", time.Now())
	noBuffer <- 1
	fmt.Println("send success: ", time.Now())
}

func noBufferReceive(noBuffer chan int) {
	fmt.Println("wait 1 second before receive")
	time.Sleep(1 * time.Second)
	// 从channel中接受数据
	<- noBuffer // 数据不使用直接丢弃
	fmt.Println("receive success: ", time.Now())
}

// channel的定义、只读、只写、双向channel演示
func define() {
	// 定义只读channel
	var onlyRead = make(<-chan int)
	// 如果是只读channel，那么不能写数据，从只读管道中读取数据时，会触发deadlock异常，因此这里的程序无法正常执行
	fmt.Println("only read channel: ", <- onlyRead)
	// 这里会编译报错：send to receive-only type <-chan int
	//onlyRead <- 20
	// 定义只写channel
	var onlyWrite = make(chan<- int)
	onlyWrite <- 10
	// 这里会编译报错，send-only type chan chan<- int
	//fmt.Println("read from only write channel: ", <- onlyWrite)
	// 定义普通channel
	var channel = make(chan int)
	channel <- 20
	close(channel)
	fmt.Println("read from channel: ", <- channel)
}