package main

import "unsafe"

import (
	"fmt"
)

var x, y int
var (
	a int
	b bool
)
var t = 1
var c, d int = 1, 2
var e, f = 123, "hello"

// 常量枚举
const (
	a2 = "ssss"
	b2 = len(a2)
	c2 = unsafe.Sizeof(a2)
)

const (
	// iota 在常量声明中，iota 的值从 0 开始递增，每次递增 1
	a4 = iota
	//b3 = iota
	//c3 = iota
	b3 // 简写
	c3
)

const (
	ii = 1 << iota
	j  = 3 << iota
	k
	l
)

func main() {
	// 定义变量
	var a string = "dadada"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var lzc = "www"
	fmt.Println(lzc)

	// 声明一个变量并初始化
	var all = "Rooo"
	fmt.Println(all)
	// 若没有初始化，则赋值为零值
	var nnn string
	fmt.Println(nnn)
	var cc bool
	fmt.Println(cc)

	// %d 表示整型数字，%s 表示字符串
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"

	// 第一种占位符方式
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
	// 第二种占位符方式
	fmt.Printf(url, stockcode, enddate)
	fmt.Println()

	// 定义变量
	//var aaaa *int
	//var aArray []int
	//var aMap map[string]int
	//var aaa chan int
	//var aF func(string) int
	//var aError error
	//fmt.Printf(aaaa, aArray, aMap, aaa, aF, aError)

	var i int
	var f float64
	var d bool
	var s string
	fmt.Printf("%v %v %v %q\n\n", i, f, d, s)

	intVal := 1
	//var intVal int = 1
	fmt.Println(intVal)

	g, h := 123, "hello"
	fmt.Println(x, y, a, b, c, d, e, f, g, h)

	old_a := 10
	new_a := old_a
	fmt.Println(old_a, new_a)
	new_a += 10
	fmt.Println(old_a, new_a)

	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a1, b1, c1 = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d\n", area)
	println()
	println(a1, b1, c1)

	println(a2, b2, c2)

	println(a4, b3, c3)
	println(ii, j, k, l)

	a4 := 21
	b4 := 22
	var c4 int
	c4 = a4 + b4
	fmt.Printf("c的值为 %d\n", c4)
	c4 = a4 - b4
	fmt.Printf("c的值为 %d\n", c4)
	c4 = a4 * b4
	fmt.Printf("c的值为 %d\n", c4)
	c4 = a4 / b4
	fmt.Printf("c的值为 %d\n", c4)
	c4 = a4 % b4
	fmt.Printf("c的值为 %d\n", c4)
	a4++
	fmt.Printf("a的值为 %d\n", a4)
	a4 = 21
	a4--
	fmt.Printf("a的值为 %d\n", a4)

	var a5 int = 21
	var b5 int = 10

	if a5 == b5 {
		fmt.Printf("第一行 - a 等于 b\n")
	} else {
		fmt.Printf("第一行 - a 不等于 b\n")
	}
	if a5 < b5 {
		fmt.Printf("第二行 - a 小于 b\n")
	} else {
		fmt.Printf("第二行 - a 不小于 b\n")
	}

	if a5 > b5 {
		fmt.Printf("第三行 - a 大于 b\n")
	} else {
		fmt.Printf("第三行 - a 不大于 b\n")
	}
	/* Lets change value of a and b */
	a5 = 5
	b5 = 20
	if a5 <= b5 {
		fmt.Printf("第四行 - a 小于等于 b\n")
	}
	if b5 >= a5 {
		fmt.Printf("第五行 - b 大于等于 a\n")
	}

	var a6 bool = true
	var b6 bool = false
	if a6 && b6 {
		fmt.Printf("第一行 - 条件为 true\n")
	}
	if a6 || b6 {
		fmt.Printf("第二行 - 条件为 true\n")
	}
	/* 修改 a 和 b 的值 */
	a6 = false
	b6 = true
	if a6 && b6 {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if !(a6 && b6) {
		fmt.Printf("第四行 - 条件为 true\n")
	}

	// 位运算符
	var a8 uint = 60 /* 60 = 0011 1100 */
	var b8 uint = 13 /* 13 = 0000 1101 */
	var c8 uint = 0

	// 按位与 有0则0
	c8 = a8 & b8 /* 12 = 0000 1100 */
	fmt.Printf("第一行 - c 的值为 %d\n", c8)
	// 按位或 有1则1
	c8 = a8 | b8 /* 61 = 0011 1101 */
	fmt.Printf("第二行 - c 的值为 %d\n", c8)
	// 异或运算 相异为1
	c8 = a8 ^ b8 /* 49 = 0011 0001 */
	fmt.Printf("第三行 - c 的值为 %d\n", c8)

	c8 = a8 << 2 /* 240 = 1111 0000 */
	fmt.Printf("第四行 - c 的值为 %d\n", c8)

	c8 = a8 >> 2 /* 15 = 0000 1111 */
	fmt.Printf("第五行 - c 的值为 %d\n", c8)

	// 赋值运算符
	var a7 int = 21
	var c7 int

	c7 = a7
	fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c7)

	c7 += a7
	fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c7)

	c7 -= a7
	fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c7)

	c7 *= a7
	fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c7)

	c7 /= a7
	fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c7)

	c7 = 200

	c7 <<= 2
	fmt.Printf("第 6行  - <<= 运算符实例，c 值为 = %d\n", c7)

	c7 >>= 2
	fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c7)

	c &= 2
	fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c7)

	c ^= 2
	fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c7)

	c |= 2
	fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c7)

	var a9 int = 4
	var b9 int32
	var c9 float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a9)
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b9)
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c9)
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", ptr)

	/*  & 和 * 运算符实例 */
	ptr = &a9 /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("a 的值为  %d\n", a9)
	fmt.Printf("*ptr 为 %d\n", *ptr)

	// 运算符优先级
	var a10 int = 20
	var b10 int = 10
	var c10 int = 15
	var d10 int = 5
	var e10 int

	e10 = (a10 + b10) * c10 / d10 // ( 30 * 15 ) / 5
	fmt.Printf("(a10 + b10) * c10 / d10 的值为 : %d10\n", e10)

	e10 = ((a10 + b10) * c10) / d10 // (30 * 15 ) / 5
	fmt.Printf("((a10 + b10) * c10) / d10 的值为  : %d10\n", e10)

	e10 = (a10 + b10) * (c10 / d10) // (30) * (15/5)
	fmt.Printf("(a10 + b10) * (c10 / d10) 的值为  : %d10\n", e10)

	e10 = a10 + (b10*c10)/d10 //  20 + (150/5)
	fmt.Printf("a10 + (b10 * c10) / d10 的值为  : %d10\n", e10)

}
