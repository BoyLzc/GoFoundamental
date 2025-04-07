package main

import "fmt"

/*
1. 数组：是同一种数据类型的固定长度的序列。
2. 数组定义：var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。
3. 长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。
4. 数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1
for i := 0; i < len(a); i++ {
}
for index, v := range a {
}
5. 访问越界，如果下标在数组合法范围之外，则触发访问越界，会panic
6. 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
7.支持 "=="、"!=" 操作符，因为内存总是被初始化过的。
8.指针数组 [n]*T，数组指针 *[n]T。
*/
func scanArray(a [5]int) {
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
}

func printArray(a [5]int) {
	for index, v := range a {
		fmt.Println(index, v)
	}
}

func test(x [2]int) {
	fmt.Printf("x :%p\n", &x)
	x[1] = 1000
}

func findIndex(array [5]int, target int) {
	for i := 0; i < len(array); i++ {
		d := target - array[i]
		for j := i + 1; j < len(array); j++ {
			if d == array[j] {
				fmt.Println("found")
				fmt.Println(i, j)
			}
		}
	}
	fmt.Println("not found")
}

func main() {
	var a [5]int
	scanArray(a)
	fmt.Println(a)
	printArray(a)
	// 数组和切片声明方式不同，切片[] 数组必须声明长度或使用[...]
	var arr0 [5]int = [5]int{1, 2, 3}
	var arr1 = [5]int{1, 2, 3, 4, 5}
	var arr2 = [...]int{1, 2, 3, 4, 5, 6}
	var str = [5]string{3: "hello world", 4: "tom"}
	fmt.Println(arr0, arr1, arr2, str)

	var arr3 [5][3]int
	var arr4 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}
	c := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	b := [4][2]int{{1, 1}, {2, 2}, {3, 3}}
	d := [...][2]int{{2, 2}, {3, 3}}
	l := []int{2, 3, 4}
	fmt.Println(l)
	l = append(l, 1)
	fmt.Println(l)

	fmt.Println(arr3, arr4, c, b, d)

	e := [2]int{}
	fmt.Printf("a: %p\n", &e)

	test(e)
	fmt.Println(e)

	var f [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

	for k1, v1 := range f {
		fmt.Println(v1) // 一维数组
		for k2, v2 := range v1 {
			fmt.Printf("(%d, %d)=%d ", k1, k2, v2)
		}
		fmt.Println()
	}

	array1 := [5]int{1, 2, 3, 4, 5}
	findIndex(array1, 6)

}
