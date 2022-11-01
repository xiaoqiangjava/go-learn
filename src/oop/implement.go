package main

import "fmt"

/*
go中的interface类型，是一组方法签名，当某一个类型为接口中的所有方法提供定义时，它被称为实现接口
接口中指定了类型应该具有的方法，类型决定了如何实现这些方法
接口的定义：
type interface_name interface {
	method_name(args arg_type) [return_type]
}
 */
func main() {
	// 由于IPhone实现了Phone接口的所有方法，因此IPone是对象可以直接赋值给Phone类型
	var iphone Phone = IPhone{name: "iphone13", price: 8988}
	var huawei Phone = Huawei{name: "huawei-meta30", price: 9999}
	iphone.call()
	iphone.sendMsg("iphone send msg!")
	huawei.call()
	huawei.sendMsg("huawei send msg!")

	// 空接口：没有任何方法的接口就是空接口，因此任何类型都实现了空接口，这个类似于Java中的Object类型，空接口可以存储任何类型的值
	var obj Object = IPhone{name: "iphone14", price: 9988}
	// 空接口中取值需要使用类型断言，这里的obj.(IPhone)就是告诉程序我们的obj是一个IPhone类型
	var name = obj.(IPhone).name
	fmt.Println("name: ", name)
	// 空接口可以不用命名，直接写为interface{}, 因为空接口是一种数据类型
	var any interface{} = "1"
	fmt.Println("any: ", any)
}

// 定义接口
type Phone interface {
	call()
	sendMsg(msg string) bool // 有参数有返回值
}

type Object interface {
}

// 定义一个IPhone类型
type IPhone struct {
	name string
	price float64
}

// 定义一个Huawei类型
type Huawei struct {
	name string
	price float64
}

// 实现接口：需要实现接口中的所有方法，go语言中的实现就是将方法与类型进行一个绑定
// IPhone实现call方法
func (iphone IPhone) call() {
	fmt.Printf("%s打电话\n", iphone.name)
}

// IPhone实现sendMsg方法
func (iphone IPhone) sendMsg(msg string) bool {
	fmt.Printf("%s发短信, 内容为: %s\n", iphone.name, msg)
	return true
}

// Huawei实现call方法

func (huawei Huawei) call() {
	fmt.Printf("%s打电话\n", huawei.name)
}

// Huawei实现sendMsg方法
func (huawei Huawei) sendMsg(msg string) bool {
	fmt.Printf("%s发短信, 内容为: %s\n", huawei.name, msg)
	return true
}