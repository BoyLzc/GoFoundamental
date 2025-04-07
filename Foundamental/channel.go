package main

import "fmt"

// 信道就是一个管道，连接多个goroutine程序 是一种队列式的数据结构，遵循先进先出的规则
func main() {

	// 定义信道
	//pipline := make(chan int)

	/*	// 往信道中发送数据
		go func() {
			pipline <- 200
		}()
		// 从信道中取出数据，并赋值给mydata
		mydata := <-pipline*/
	//x, ok := <-pipline
	//fmt.Println(x, ok)

	pipline2 := make(chan int, 10)
	fmt.Printf("信道可缓冲 %d 个数据\n", cap(pipline2))
	pipline2 <- 1
	fmt.Printf("信道中当前有 %d 个数据", len(pipline2))
	fmt.Printf("信道中当前有 %d 个数据", len(pipline2))

}
