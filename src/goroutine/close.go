package main

import (
	"fmt"
	"time"
)

func main() {
	ctx := Context{
		write: make(chan int, 1),
		done:  make(chan int),
	}
	go testOne(ctx)
	go testTwo(ctx)
	for i := 0; i < 10; i++ {
		ctx.write <- i
	}

	close(ctx.done)
	time.Sleep(2 * time.Second)

}

func testOne(ctx Context) {
	for {
		select {
		case i := <-ctx.write:
			fmt.Println("test one read i: ", i)
		case <-ctx.done:
			fmt.Println("test one context done")
			return
		}
	}
}

func testTwo(ctx Context) {
	for {
		select {
		case i := <-ctx.write:
			fmt.Println("test two read i: ", i)
		case <-ctx.done:
			fmt.Println("test two context done")
			return
		}
	}
}

type Context struct {
	done  chan int
	write chan int
}
