package main

import "fmt"

// 函数
func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

// 函数可以返回多个值 默认情况下，go语言使用的是值传递
func swap(x, y string) (string, string) {
	return y, x
}

func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)
	return a + b
}

/* 声明全局变量 */
var g int = 20

var a3 int = 50

func main() {
	a := 100
	b := 200
	c := max(a, b)
	println(c)
	e, f := swap("123", "456")
	fmt.Println(e, f)

	/* 声明局部变量 */
	var a2, b2, c2 int
	/* 初始化参数 */
	a2 = 10
	b2 = 20
	c2 = a2 + b2
	fmt.Printf("结果： a2 = %d, b2 = %d and c2 = %d\n", a2, b2, c2)

	g = a2 + b2
	fmt.Printf("g的值为 %d\n", g)

	/* 声明局部变量 */
	var g int = 10
	fmt.Printf("g的值为 %d\n", g)

	/* main 函数中声明局部变量 */
	var a3 int = 10
	var b3 int = 20
	var c3 int = 0

	fmt.Printf("main()函数中 a = %d\n", a3)
	c3 = sum(a3, b3)
	fmt.Printf("main()函数中 c = %d\n", c3)
}
