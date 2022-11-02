package main

import "fmt"

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
make(chan int, 10) // 带缓冲区
make(chan int) // 不带缓冲区
注意：
channel中的数据必须是成对出现的，有发送，有读取，如果直接收，没有发送或者只发送没有接收，那么所有的goroutine无法运行，程序将会以deadlock结束
因此在定义channel时，如果定义了只读、只写chan，是没有意义的，程序会以deadlock结束
channel主要用来协程之间通信，因此发送方和接收方一般都是不同的协程
 */
func main() {
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
