package main

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"time"
)

// go build test.go  编译go源文件
// go run test.go  编译并运行go源文件
// gofmt 格式化源文件，-w参数会修改源文件
func main() {
	start := time.Now().UnixMilli()
	defer fmt.Println("time", time.Now().UnixMilli()-start)

	fmt.Println("len: ", len("同意"))

	time.Sleep(10 * time.Second)

	fmt.Println("time duration: ", time.Now().Unix()-start)

	fmt.Println("start goroutine num: ", runtime.NumGoroutine())

	fmt.Println("end goroutine num: ", runtime.NumGoroutine())
	defer fmt.Println("main done")
}

func testLock(num int, lock *Lock) {
	lock.lock.Lock()
	if num%2 == 0 {
	}
	lock.lock.Unlock()
	fmt.Println("lock is not unlock")
}

type Lock struct {
	lock sync.Mutex
}

func async(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("async panic: ", r)
			}
		}()
		fn()
	}()
}

func contains(e interface{}, list interface{}) bool {
	tp := reflect.TypeOf(list)
	fmt.Println(tp.Elem())
	val := reflect.ValueOf(list)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(e, val.Index(i).Interface()) {
				return true
			}
		}
	}

	return false
}
