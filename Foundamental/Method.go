package main

import "fmt"

type Test struct{}

type Data struct {
	x int
}

func (self Data) ValueTest() { // func ValueTest(self Data);
	fmt.Printf("Value: %p\n", &self)
}

func (self *Data) PointerTest() { // func PointerTest(self *Data);
	fmt.Printf("Pointer: %p\n", self)
}

// 无参数、无返回值
func (t Test) method0() {

}

// 单参数、无返回值
func (t Test) method1(i int) {

}

// 多参数、无返回值
func (t Test) method2(x, y int) {

}

// 无参数、单返回值
func (t Test) method3() (i int) {
	return
}

// 多参数、多返回值
func (t Test) method4(x, y int) (z int, err error) {
	return
}

// 无参数、无返回值
func (t *Test) method5() {

}

// 单参数、无返回值
func (t *Test) method6(i int) {

}

// 多参数、无返回值
func (t *Test) method7(x, y int) {

}

// 无参数、单返回值
func (t *Test) method8() (i int) {
	return
}

// 多参数、多返回值
func (t *Test) method9(x, y int) (z int, err error) {
	return
}

// 结构体
type User struct {
	Name  string
	Email string
}

// 方法
func (u User) Notify() {
	u.Name = "123"
	u.Email = "123@123.com"
	//fmt.Printf("%v : %v \n", u.Name, u.Email)
}

func (u *User) Notify2() {
	u.Name = "123"
	u.Email = "123@123.com"
	//fmt.Printf("%v : %v \n", u.Name, u.Email)
}

// 1.普通函数
// 接收值类型参数的函数
func valueIntTest(a int) int {
	return a + 10
}

// 接收指针类型参数的函数
func pointerIntTest(a *int) int {
	return *a + 10
}

func structTestValue() {
	a := 2
	fmt.Println("valueIntTest:", valueIntTest(a))
	//函数的参数为值类型，则不能直接将指针作为参数传递
	//fmt.Println("valueIntTest:", valueIntTest(&a))
	//compile error: cannot use &a (type *int) as type int in function argument

	b := 5
	fmt.Println("pointerIntTest:", pointerIntTest(&b))
	//同样，当函数的参数为指针类型时，也不能直接将值类型作为参数传递
	//fmt.Println("pointerIntTest:", pointerIntTest(&b))
	//compile error:cannot use b (type int) as type *int in function argument
}

// 2.方法
type PersonD struct {
	id   int
	name string
}

// 接收者为值类型
func (p PersonD) valueShowName() {
	fmt.Println(p.name)
}

// 接收者为指针类型
func (p *PersonD) pointShowName() {
	fmt.Println(p.name)
}

func structTestFunc() {
	//值类型调用方法
	personValue := PersonD{101, "hello world"}
	personValue.valueShowName()
	personValue.pointShowName()

	//指针类型调用方法
	personPointer := &PersonD{102, "hello golang"}
	personPointer.valueShowName()
	personPointer.pointShowName()

	//与普通函数不同，接收者为指针类型和值类型的方法，指针类型和值类型的变量均可相互调用
}

func main() {
	// demo1
	// 值类型调用方法
	/*	u1 := User{"golang", "golang@golang.com"}
		u1.Notify()
		fmt.Printf("%v : %v \n", u1.Name, u1.Email)
		u1.Notify2()
		fmt.Printf("%v : %v \n", u1.Name, u1.Email)

		// 指针类型调用方法
		u2 := User{"go", "go@go.com"}
		u3 := &u2
		//解释： 首先我们定义了一个叫做 User 的结构体类型，
		//然后定义了一个该类型的方法叫做 Notify，该方法的接受者是一个 User 类型的值。
		//要调用 Notify 方法我们需要一个 User 类型的值或者指针。
		//在这个例子中当我们使用指针时，Go 调整和解引用指针使得调用可以被执行。
		//注意，当方法定义的接受者不是一个指针时，该方法操作对应接受者的值的副本(意思就是即使你使用了指针调用函数，但是函数的接受者是值类型，所以函数内部操作还是对副本的操作，而不是指针操作。
		//注意：当方法定义的接受者是指针时，即使用值类型调用那么函数内部也是对指针的操作。
		u3.Notify()
		fmt.Printf("%v : %v \n", u2.Name, u2.Email)
		u3.Notify2()
		fmt.Printf("%v : %v \n", u2.Name, u2.Email)*/

	// demo2
	/*	d := Data{} //结构体变量
		p := &d     // 变量地址

		fmt.Printf("Data: %p\n", p)

		d.ValueTest()   // ValueTest(d)
		d.PointerTest() // PointerTest(&d)

		p.ValueTest()   // ValueTest(*p)
		p.PointerTest() // PointerTest(p)*/

	// demo3
	structTestValue()
	structTestFunc()
}
