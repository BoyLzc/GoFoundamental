package main

import "fmt"

func main() {
	s := "abc"
	// 忽略 2nd value，支持 string/array/slice/map。
	for i := range s {
		println(s[i])
	}
	// 忽略 index。
	for _, c := range s {
		println(c)
	}
	// 忽略全部返回值，仅迭代。
	for range s {

	}

	m := map[string]int{"a": 1, "b": 2}
	// 返回 (key, value)。
	for k, v := range m {
		println(k, v)
	}

	a := [3]int{0, 1, 2}

	for i, v := range a { // index、value 都是从复制品中取出。

		if i == 0 { // 在修改前，我们先修改原数组。
			a[1], a[2] = 999, 999
			fmt.Println(a) // 确认修改有效，输出 [0, 999, 999]。
		}

		a[i] = v + 100 // 使用复制品中取出的 value 修改原数组。

	}

	fmt.Println(a) // 输出 [100, 101, 102]。

	b := []int{1, 2, 3}
	for i, v := range b { // index、value 都是从复制品中取出。

		if i == 0 { // 在修改前，我们先修改原数组。
			b = b[:]
			// 会改变 i ==2 时的 v
			b[2] = 100
		}

		b[i] = v + 100 // 使用复制品中取出的 value 修改原数组。
	}
	fmt.Println(b)
}
