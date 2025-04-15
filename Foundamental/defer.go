package main

import "fmt"

/*func myfunc() {
	fmt.Println("B")
}*/

var name string = "go"

func myfunc() string {
	defer func() {
		name = "python"
	}()

	fmt.Printf("myfunc 函数里的name：%s\n", name)
	return name
}

func main() {
	// demo 1
	/*	defer myfunc() // 延时调用 defet xxx() 实现将 xxx函数延迟到当前函数体执行完再执行
		fmt.Println("A")*/

	// demo2
	/*	name := "go"
		defer fmt.Println(name) // 输出: go

		name = "python"
		fmt.Println(name) // 输出: python*/

	// demo3
	/*	name := "go"
		defer func() { // defer 后跟匿名函数，情况会有所不同
			fmt.Println(name) // 输出: python
		}()
		name = "python"
		fmt.Println(name) // 输出: python*/
	// demo4
	myname := myfunc()
	fmt.Printf("main 函数里的name: %s\n", name)
	// defer 在 return 后调用
	fmt.Println("main 函数里的myname: ", myname)

	//

}
