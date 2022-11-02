package main

import (
	"fmt"
	"strings"
	"time"
)

// 定义一个LogProcess的结构体
type LogProcess struct {
	// 文件路径
	Path string
	// 数据库配置
	DbConfig struct{
		username string
		password string
		url string
	}
	// 读取文件的通道
	ReadChannel chan string
	// 写入文件的通道
	WriteChannel chan string
}

// 读取文件
func (logProcess *LogProcess) ReadFromFile() {
	// 读取文件
	fmt.Println("start read file, path: ", logProcess.Path)
	line := "message"
	fmt.Println("读取到的文件内容为: ", line)
	// 将读取到的内容发送到channel，供解析的协程使用
	logProcess.ReadChannel <- line
	fmt.Println("end read file.")
}

// 解析文件
func (logProcess *LogProcess) Process() {
	// 解析文件
	fmt.Println("start process file content.")
	// 从通道获取文件内容
	content := <- logProcess.ReadChannel
	// 这里讲文件内容修改为大写，然后输出到写通道
	logProcess.WriteChannel <- strings.ToUpper(content)
	fmt.Println("end process file content.")
}

// 写入数据库
func (logProcess *LogProcess) WriteToInfluxDB() {
	// 数据入库
	fmt.Println("start write content to influxDB.")
	// 从通道获取数据
	content := <- logProcess.WriteChannel
	fmt.Println("获取到的数据内容为: ", content)
	fmt.Println("end write content to influxDB")
}

// main
func main() {
	// 创建变量
	logProcess := &LogProcess{
		Path: "/temp/access.log",
		DbConfig: struct {
			username string
			password string
			url      string
		}{username: "root", password: "root", url: "localhost:3000"},
		ReadChannel: make(chan string),
		WriteChannel: make(chan string),  // 注意这里如果把}写到下一行，需要以逗号结尾，这为未来添加或者移除属性带来方便
	}
	// 读取文件
	go logProcess.ReadFromFile()
	// 解析文件
	go logProcess.Process()
	// 存储
	go logProcess.WriteToInfluxDB()

	// 让主程序休眠
	time.Sleep(1 * time.Second)
}