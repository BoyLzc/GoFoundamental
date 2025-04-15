package main

import (
	"fmt"
	"sync"
	"time"
	//"time"
)

// demo1
/*func add(count *int, wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		*count = *count + 1
	}
	wg.Done()
}*/

var x int64
var wg sync.WaitGroup
var lock sync.Mutex     // 互斥锁
var rwlock sync.RWMutex // 读写锁

func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
}

// 互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源。
// Go语言中使用sync包的Mutex类型来实现互斥锁。 使用互斥锁来修复上面代码的问题：
func add2() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 解锁
	}
	wg.Done()
}

func add3() {
	for i := 0; i < 5000; i++ {
		fmt.Println("我是", "执行的")
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 解锁
		fmt.Println(x)
	}
	fmt.Println("我是", "执行完毕***********")
}

func write() {
	// lock.Lock()   // 加互斥锁
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock.Unlock()                   // 解写锁
	// lock.Unlock()                     // 解互斥锁
	wg.Done()
}

func read() {
	// lock.Lock()                  // 加互斥锁
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	// lock.Unlock()                // 解互斥锁
	wg.Done()
}

/*
Add：初始值为0，你传入的值会往计数器上加，这里直接传入你子协程的数量

Done：当某个子协程完成后，可调用此方法，会从计数器上减一，通常可以使用 defer 来调用。

Wait：阻塞当前协程，直到实例里的计数器归零。
*/
func main() {
	// demo1
	/*	var wg sync.WaitGroup
		count := 0
		wg.Add(3)
		go add(&count, &wg)
		go add(&count, &wg)
		go add(&count, &wg)

		wg.Wait()
		fmt.Println("count 的值为：", count)*/

	// demo2
	// 开启了两个goroutine去累加变量x的值，
	// 这两个goroutine在访问和修改x变量的时候就会存在数据竞争，导致最后的结果与期待的不符。
	/*	go add()
		go add()
		time.Sleep(5 * time.Second)
		fmt.Println(x)*/

	// demo3
	// 互斥锁是一种常用的控制共享资源访问的方法，
	// 它能够保证同时只有一个goroutine可以访问共享资源。
	// 使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁；
	// 当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同时等待一个锁时，唤醒的策略是随机的。
	/*	wg.Add(2)
		go add2()
		go add2()
		wg.Wait()
		fmt.Println(x)*/

	// demo4
	// 互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，
	// 当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下使用读写锁是更好的一种选择。
	// 读写锁在Go语言中使用sync包中的RWMutex类型。
	//
	// 读写锁分为两种：读锁和写锁。
	// 当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
	// 当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
