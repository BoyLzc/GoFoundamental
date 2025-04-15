package main

import "fmt"

func myFunction(param []float32) {
	fmt.Println(param)
}

func InitialArray(param [10]int) [10]int {
	var i int
	for i = 0; i < 10; i++ {
		param[i] = i + 100
	}
	return param
}

func PrintArray(param [10]int) {
	var j int
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, param[j])
	}
}

func getAverage(arr []int, size int) float32 {
	var i int
	var avg, sum float32

	for i = 0; i < size; i++ {
		sum += float32(arr[i])
	}

	avg = sum / float32(size)

	return avg
}

// 函数接受一个数组作为参数
func modifyArray(arr [5]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
}

// 函数接受一个数组的指针作为参数
func modifyArrayWithPointer(arr *[5]int) {
	fmt.Println(arr)
	fmt.Println(*arr)
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = (*arr)[i] * 2
	}
}

func main() {
	var numbers1 [10]float32
	fmt.Println(numbers1)
	var numbers2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers2)
	numbers3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers3)

	balance1 := [...]float32{}
	fmt.Println(balance1)
	balance2 := []float32{1: 2.0, 3: 7.0}
	fmt.Println(balance2)

	var a float32 = balance2[1]
	fmt.Println(a)

	myFunction(balance2)

	balance3 := [10]int{}
	fmt.Println(InitialArray(balance3))

	var balance4 [10]int
	fmt.Println(InitialArray(balance4))

	PrintArray(balance4)

	/* 数组长度为 5 */
	var balance5 = []int{1000, 2, 3, 17, 50}
	var avg float32

	/* 数组作为参数传递给函数 */
	avg = getAverage(balance5, 5)

	/* 输出返回的平均值 */
	fmt.Printf("平均值为: %f ", avg)

	// 创建一个包含5个元素的整数数组
	myArray := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Original Array:", myArray)

	// 传递数组给函数，但不会修改原始数组的值
	modifyArray(myArray)
	fmt.Println("Array after modifyArray:", myArray)

	// 传递数组的指针给函数，可以修改原始数组的值
	modifyArrayWithPointer(&myArray)
	fmt.Println("Array after modifyArrayWithPointer:", myArray)

}
