package main

import "fmt"

// 使用 ...类型，表示一个元素为int类型的切片
func Sum(args ...int) int {
	var sum int
	for _, v := range args {
		sum += v
	}
	return sum
}

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func double(a int) (b int) {
	// 不能使用 := ,因为在返回值哪里已经声明了为int
	b = a * 2
	// 不需要指明写回哪个变量,在返回值类型那里已经指定了
	return
}

// 第二个参数为函数
func visit(list []int, f func(int)) {
	for _, v := range list {
		// 执行回调函数
		f(v)
	}
}

// 函数类型
// Greeting function types
type Greeting func(name string) string

func (g Greeting) say(n string) {
	fmt.Println(g(n))
}

func english(name string) string {
	return "Hello, " + name
}

func main() {
	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	MyPrintf(v1, v2, v3, v4)

	fmt.Println(Sum(1, 2, 3))

	fmt.Println(double(10))

	func(data int) {
		fmt.Println("hello", data)
	}(100)

	// 使用匿名函数直接做为参数
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println(v)
	})

	greet := Greeting(english)
	greet.say("World")

}
