package main

import "fmt"

/*
1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
2. 切片的长度可以改变，因此，切片是一个可变的数组。
3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
6. 如果 slice == nil，那么 len、cap 结果都等于 0。
*/

var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var slice0 []int = arr[2:8]
var slice1 []int = arr[0:6]        //可以简写为 var slice []int = arr[:end]
var slice2 []int = arr[5:10]       //可以简写为 var slice[]int = arr[start:]
var slice3 []int = arr[0:len(arr)] //var slice []int = arr[:]
var slice4 = arr[:len(arr)-1]

func main() {
	// 数组和切片声明方式不同，切片[] 数组必须声明长度或使用[...]
	var s1 []int
	if s1 == nil {
		fmt.Println("numpy")
	} else {
		fmt.Println("not numpy")
	}

	s2 := []int{}
	// 3.make()
	var s3 []int = make([]int, 10)
	fmt.Println(s1, s2, s3)
	// 4.初始化赋值 长度2 容量3
	var s4 []int = make([]int, 2, 3)
	fmt.Println(s4)
	s5 := []int{1, 2, 3}
	fmt.Println(s5)
	// 5.从数组切片
	arr2 := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	// 前包后不包 左闭右开
	s6 = arr2[1:4]
	fmt.Println(s6)

	fmt.Printf("全局变量：arr %v\n", arr)
	fmt.Printf("全局变量：slice0 %v\n", slice0)
	fmt.Printf("全局变量：slice1 %v\n", slice1)
	fmt.Printf("全局变量：slice2 %v\n", slice2)
	fmt.Printf("全局变量：slice3 %v\n", slice3)
	fmt.Printf("全局变量：slice4 %v\n", slice4)
	fmt.Printf("-----------------------------------\n")
	arr3 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	slice5 := arr[2:8]
	slice6 := arr[0:6]         //可以简写为 slice := arr[:end]
	slice7 := arr[5:10]        //可以简写为 slice := arr[start:]
	slice8 := arr[0:len(arr)]  //slice := arr[:]
	slice9 := arr[:len(arr)-1] //去掉切片的最后一个元素
	fmt.Printf("局部变量： arr2 %v\n", arr3)
	fmt.Printf("局部变量： slice5 %v\n", slice5)
	fmt.Printf("局部变量： slice6 %v\n", slice6)
	fmt.Printf("局部变量： slice7 %v\n", slice7)
	fmt.Printf("局部变量： slice8 %v\n", slice8)
	fmt.Printf("局部变量： slice9 %v\n", slice9) //去掉切片的最后一个元素

	data := [...]int{0, 1, 2, 3, 4, 5}

	s := data[2:4]
	s[0] += 100
	s[1] += 200
	data[3] += 300
	// 切片改变，原数组也会改变
	fmt.Println(s)
	fmt.Println(data)

	s2 = []int{0, 1, 2, 3}
	p := &s2[2] // *int, 获取底层数组元素指针。
	*p += 100

	fmt.Println(s2)

	data2 := [][]int{
		[]int{1, 2, 3},
		[]int{100, 200},
		[]int{11, 22, 33, 44},
	}
	fmt.Println(data2)

	// 结构体数组
	d := [5]struct {
		x int
	}{}

	ss := d[:]

	d[1].x = 10
	ss[2].x = 20

	fmt.Println(d)
	fmt.Println(ss)
	fmt.Printf("%p, %p\n", &d, &d[0])
	fmt.Printf("%p, %p\n", &d, &d[1])
	fmt.Printf("%p, %p\n", &d, &d[2])
	fmt.Printf("%p, %p\n", &d, &d[3])

	s20 := []int{1, 2, 3: 100} // 通过初始化表达式构造，可使用索引号
	fmt.Println(s20, len(s20), cap(s20))

	s21 := make([]int, 6, 8) // 通过初始化表达式构造，可使用索引号
	fmt.Println(s21, len(s21), cap(s21))

	s22 := make([]int, 6) // 通过初始化表达式构造，可使用索引号
	fmt.Println(s22, len(s22), cap(s22))

	var lll = []int{1, 2, 3}
	fmt.Printf("%v\n", lll)

	var xxx = []int{4, 5, 6}
	fmt.Printf("%v\n", xxx)

	c := append(lll, xxx...)
	fmt.Printf("%v\n", c)

	ssssss := make([]int, 0, 5)
	fmt.Printf("%v\n", ssssss)
	fmt.Printf("%p\n", &ssssss)

	ssssss2 := append(ssssss, 1)
	fmt.Printf("%p\n", &ssssss2)

	data3 := [...]int{1, 2, 10: 0}
	fmt.Println(data3)
	// 设定长度:2 容量：3
	ssss := data3[:2:3]
	fmt.Println(ssss)
	fmt.Println(&ssss[0], &data3[0])

	ssss = append(ssss, 100, 200)
	// 在append之后，slice长度超过cap限制，即会重新分配底层数组，即，与原始数组地址不一样
	// 建议一次性分配足够大的空间，以减少内存分配和数据复制开销
	fmt.Println(&ssss[0], &data3[0])
	//fmt.Println(data2)

	ssss = make([]int, 0, 1)
	cc := cap(ssss)
	for i := 0; i < 100; i++ {
		ssss = append(ssss, i)
		// 通常以 两倍容量重新分配底层数组
		fmt.Println(len(ssss), cap(ssss))
		if cc != cap(ssss) {
			fmt.Println("cap changed")
			cc = cap(ssss)
		}
	}

	slicea := []int{1, 2, 3, 4, 6, 7}
	sliceb := []int{0, 2, 3, 4, 5, 6}
	// 将 slicea的值 赋值给 sliceb
	copy(sliceb, slicea)
	fmt.Println(slicea)
	fmt.Println(sliceb)

	data4 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := data4[:]
	for index, value := range slice {
		fmt.Println(index, value)
	}

	// 字符串本身不可变
	str := "Hello world"
	str2 := "我是你爹"
	/*	sss := str[:]
		fmt.Println(sss)
		fmt.Println(str)*/
	// 将其转换为 ascil值
	sss := []byte(str)
	fmt.Println(sss)
	sss[6] = 'G'
	sss = append(sss, '1')
	str = string(sss)
	fmt.Println(str)
	// 中文字符需要用[]rune(str)
	sss2 := []rune(str2)
	sss2[0] = '你'
	str2 = string(sss2)
	fmt.Println(str2)

	slice11 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice11)
	d1 := slice[6:8]

	fmt.Println(d1, len(d1), cap(d1))
	d2 := slice[:6:8]
	fmt.Println(d2, len(d2), cap(d2))

}
