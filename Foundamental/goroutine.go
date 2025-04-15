package main

import (
	"fmt"
	"time"
)

func mytest() {
	time.Sleep(time.Second)
	fmt.Println("hello, go")
}
func mygo(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("In goroutine %s\n", name)
		// 为了避免第一个协程执行过快，观察不到并发的效果，加个休眠
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	// 启动一个协程
	go mytest()
	fmt.Println("hello, world")

	go mygo("协程1号") // 第一个协程
	go mygo("协程2号") // 第二个协程
	time.Sleep(time.Second)

	//time.Sleep(time.Second)
}
