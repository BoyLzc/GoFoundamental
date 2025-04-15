package main

import "fmt"

func main() {
	//a := 10
	//b := &a
	//fmt.Printf("a：%d ptr : %p\n", a, &a)
	//fmt.Printf("b：%p type : %T\n", b, b)
	//fmt.Println(&b)

	a := 10
	b := &a
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)

	// 空指针
	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是%p\n", p)
	if p != nil {
		fmt.Println("p不为空")
	} else {
		fmt.Println("p为空")
	}

	aa := new(int)
	bb := new(bool)
	fmt.Printf("%T\n", aa)
	fmt.Printf("%T\n", bb)
	fmt.Println(*aa)
	fmt.Println(*bb)
	// 对于引用类型的变量，在使用时不仅要声明还要分配内存
	// new 函数 和 make 函数 主要用于分配内存
	// 声明
	var aaa *int
	// 打印为空 因为还未分配内存
	fmt.Println(aaa)
	// 分配内存
	aaa = new(int)
	fmt.Println(aaa)
	*aaa = 10
	fmt.Println(*aaa)

	var lzc int
	fmt.Println(lzc)
	fmt.Println(&lzc)
	fmt.Println()
	lzc2 := new(int)
	fmt.Println(lzc2)
	fmt.Println(&lzc2)
	fmt.Println()

	// make 函数 作用对象 slice map channel
	var lzc3 map[string]int
	fmt.Println(lzc3)
	fmt.Println(&lzc3)
	fmt.Printf("%T\n", lzc3)

	lzc3 = make(map[string]int, 10)
	fmt.Println(lzc3)
	fmt.Println(&lzc3)

	lzc3["测试"] = 100
	fmt.Println(lzc3)

}
