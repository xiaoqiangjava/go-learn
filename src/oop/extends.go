package main

import "fmt"

/*
使用struct嵌套可以实现继承
 */

// 定义一个Foo类型
type Foo struct {
	name string
}

// 给Foo类型定义一个方法
func (foo Foo) echo() {
	fmt.Println("调用了Foo类型的echo方法, foo: ", foo.name)
}

// 定义一个Bar类型，继承Foo，通过匿名字段的方式，让Bar类型继承了Foo类型的所有方法和属性
type Bar struct {
	Foo // 这个是匿名字段，类型为Foo的匿名字段
	num int
}

func main() {
	// 创建Bar类型的变量, 这里虽然是继承但是不能把Bar类型的变量赋值给Foo，要想实现多肽，需要使用interface来实现
	var bar = Bar{num: 1}
	// 由于Bar继承了Foo，因此具有Foo的所有属性和方法
	bar.name = "bar"
	bar.echo()
}