package main

import "fmt"

func main() {
	// demo1
	/*		goto flag
		fmt.Println("B")
	flag:
		fmt.Println("A")*/

	// demo2
	/*		i := 1
	flag:
		if i <= 5 {
			fmt.Println(i)
			i++
			goto flag
		}*/

	// demo3
	/*	i := 1
		flag:
			for i <= 10 {
				// i 为 奇数才会进入循环
				if i%2 == 1 {
					i++
					goto flag
				}
				// 只会打印偶数
				fmt.Println(i)
				i++
			}*/

	// demo4
	/*	fmt.Println("start")
			goto flag // goto 语句之后不能有 变量声明
			var say = "hello oldboy"
			fmt.Println(say)
		flag:
			fmt.Println("end")*/
	fmt.Println("lzzc")

}
