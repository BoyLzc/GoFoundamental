package main

import (
	"fmt"
<<<<<<< HEAD
<<<<<<< HEAD
	"sync"
=======
	//"time"
>>>>>>> f208e27996a553baa224bac3e753deed5b11c91f
=======
	//"time"
>>>>>>> c407bbab280455c56d69b21dbdd364016525049b
)

// 定义只写信道类型
type Sender = chan<- int

// 定义只读信道类型
type Receiver = <-chan int

func fibonacci(mychan chan int) {
	n := cap(mychan)
	x, y := 1, 1
	for i := 0; i < n; i++ {
		mychan <- x
		x, y = y, x+y
		//fmt.Println(x, y)
	}
	// 记得 close 信道
	// 不然主函数中遍历完并不会结束，而是会阻塞。
	close(mychan)
}

// 由于 x=x+1 不是原子操作
// 所以应避免多个协程对x进行操作
// 使用容量为1的信道可以达到锁的效果
func increment(ch chan bool, x *int) {
	ch <- true
	*x = *x + 1
	<-ch
}

<<<<<<< HEAD
<<<<<<< HEAD
func worker(x int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("worker %d: %d\n", x, i)
	}
}

=======
>>>>>>> f208e27996a553baa224bac3e753deed5b11c91f
=======
>>>>>>> c407bbab280455c56d69b21dbdd364016525049b
// 信道就是一个管道，连接多个goroutine程序 是一种队列式的数据结构，遵循先进先出的规则
func main() {
	// demo1
	// 定义信道
	//pipline := make(chan int)

	/*		// 往信道中发送数据
		go func() {
			pipline <- 200
		}()
		// 从信道中取出数据，并赋值给mydata
		mydata := <-pipline
	//x, ok := <-pipline
	//fmt.Println(x, ok)*/

	// demo2
	/*	pipline2 := make(chan int, 10)
		fmt.Printf("信道可缓冲 %d 个数据\n", cap(pipline2))
		pipline2 <- 1
		fmt.Printf("信道中当前有 %d 个数据", len(pipline2))*/

	// demo3
	/*	pipline := make(chan int)
		go func() {
			num := <-pipline
			fmt.Printf("接收到的数据是: %d", num)
		}()

		go func() {
			fmt.Println("准备发送数据: 100")
			pipline <- 100
		}()

		// 主函数sleep，使得上面两个goroutine有机会执行
		time.Sleep(time.Second)*/

	// demo4
	/*	var pipline = make(chan int)

		go func() {
			var sender Sender = pipline
			fmt.Println("准备发送数据: 100")
			sender <- 100
		}()

		go func() {
			var receiver Receiver = pipline
			num := <-receiver
			fmt.Printf("接收到的数据是: %d", num)
		}()
		// 主函数sleep，使得上面两个goroutine有机会执行
		time.Sleep(time.Second)*/

	// demo5 遍历信道
	/*	pipline := make(chan int, 10)

		go fibonacci(pipline)

		for k := range pipline {
			fmt.Println(k)
		}*/

	// demo6
	// 注意要设置容量为 1 的缓冲信道
	/*	pipline := make(chan bool, 1)

		var x int
		for i := 0; i < 1000; i++ {
			go increment(pipline, &x)
		}

		// 确保所有的协程都已完成
		// 以后会介绍一种更合适的方法（Mutex），这里暂时使用sleep
		time.Sleep(time.Second)
		fmt.Println("x 的值：", x)*/

	// demo7 浅拷贝与深拷贝
	/*	aArr := [3]int{0, 1, 2}
		fmt.Printf("打印 aArr: %v \n", aArr)
		bArr := aArr
		aArr[0] = 88
		fmt.Println("将 aArr 拷贝给 bArr 后，并修改 aArr[0] = 88")
		fmt.Printf("打印 aArr: %v \n", aArr)
		fmt.Printf("打印 bArr: %v \n", bArr)

		aslice := []int{0, 1, 2}
		fmt.Printf("打印 aslice: %v \n", aslice)
		bslice := aslice
		aslice[0] = 88
		fmt.Println("将 aslice 拷贝给 bslice 后，并修改 aslice[0] = 88")
		fmt.Printf("打印 aslice: %v \n", aslice)
		fmt.Printf("打印 bslice: %v \n", bslice)*/

<<<<<<< HEAD
<<<<<<< HEAD
	// demo 8
	/*	done := make(chan bool)
		go func() {
			for i := 0; i < 5; i++ {
				fmt.Println(i)
			}
			done <- true
		}()
		// 会一直等待
		<-done*/

	// demo9
	var wg sync.WaitGroup

	/*
		Add：初始值为0，你传入的值会往计数器上加，这里直接传入你子协程的数量
		Done：当某个子协程完成后，可调用此方法，会从计数器上减一，通常可以使用 defer 来调用。
		Wait：阻塞当前协程，直到实例里的计数器归零。
	*/
	wg.Add(2)
	go worker(1, &wg)
	go worker(2, &wg)

	wg.Wait()
=======
>>>>>>> f208e27996a553baa224bac3e753deed5b11c91f
=======
>>>>>>> c407bbab280455c56d69b21dbdd364016525049b
}
