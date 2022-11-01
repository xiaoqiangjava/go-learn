package main

import "fmt"

/*
go 语言中的对象，是通过结构体来实现的，自定义一个类型通过type关键字
*/
type Teacher struct {
	Name   string
	Age    int
	School string
}

// 给类型绑定方法，达到封装的效果, 使用首字母大写实现权限的管控
func (teacher Teacher) teach() {
	fmt.Printf("%s老师正在讲课\n", teacher.Name)
}

func main() {
	// 创建类型对应的对象或者变量
	// 方式一
	var t1 Teacher
	t1.Name = "Eason"
	t1.Age = 18
	t1.School = "清华"
	fmt.Println(t1)
	fmt.Println(t1.Age)

	// 方式二
	var t2 Teacher = Teacher{}
	t2.Name = "Eason"
	t2.Age = 23
	t2.School = "北京大学"
	fmt.Println(t2)
	// 这种方式可以直接在{}中对属性赋值，按照申明的顺序复制
	var t3 Teacher = Teacher{"eason", 23, "上海大学"}
	fmt.Println(t3)

	// 方式三，使用new()函数，返回的是一个指针，但是赋值时go内部做了优化，可以使用使用指针.的方式，不需要带*
	var t4 *Teacher = new(Teacher)
	t4.Name = "Eason"
	t4.Age = 98
	(*t4).School = "复旦"
	fmt.Println(*t4)

	// 方式四
	var t5 *Teacher = &Teacher{"Eason", 18, "fudan"}
	fmt.Printf("指针: %p, 值: %v\n", t5, *t5)

	// 调用teacher的方法
	t5.teach()

}
