package main

import "fmt"

// go build test.go  编译go源文件
// go run test.go  编译并运行go源文件
// gofmt 格式化源文件，-w参数会修改源文件
func main() {
	fmt.Println("hello world!")
	fmt.Println("helo golang")
	fmt.Println("hello eason")
	var age = 19 + 23
	fmt.Println("age: ", age)
}
