package main

import "fmt"

func set_data(x int) {
	defer func() {
		// recover() 可以将捕获到的panic信息打印
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	// 故意制造数组越界，触发 panic
	var arr [10]int
	arr[x] = 88
}

func main() {
	//panic("crash")

	// recover 使用，它可以让程序在发生宕机后起死回生，但必须在defer中调用
	set_data(20)

	// 如果能执行到这句，说明panic被捕获了
	// 后续的程序能继续运行
	fmt.Println("everything is ok")
}
