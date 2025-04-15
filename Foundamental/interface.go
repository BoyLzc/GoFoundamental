package main

import (
	"fmt"
	//"unicode"
)

// import "fmt"
// 接口是一种数据类型
// Sayer 接口
/*type Sayer interface {
	Say() string
}

type Mover interface {
	move()
	//move2()
}

type Cat struct{}

func (c Cat) Say() string { return "喵喵喵" }

type Dog struct{}

func (d Dog) Say() string { return "汪汪汪" }
func (d Dog) move() {
	fmt.Println("狗会动")
}

func (d *Dog) move2() {
	fmt.Println("狗会动2")
}

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}*/

// Mover 接口 demo 6
/*type Mover interface {
	move()
}

type dog struct {
	name string
}

type car struct {
	brand string
}

// dog类型实现Mover接口
func (d dog) move() {
	fmt.Printf("%s会跑\n", d.name)
}

// car类型实现Mover接口
func (c car) move() {
	fmt.Printf("%s速度70迈\n", c.brand)
}*/

// demo 7
// 一个接口的方法，不一定需要由一个类型完全实现，
// 接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。
/*type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct {
	name string
}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}*/

// demo8 接口嵌套
type Sayer interface {
	say()
}

type Mover interface {
	move()
}

type animal interface {
	Sayer
	Mover
}

type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

func (c cat) move() {
	fmt.Println("猫会动")
}

func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}

func main() {
	//c := Cat{}
	//fmt.Println("猫:", c.Say())
	//d := Dog{}
	//fmt.Println("狗:", d.Say())

	// demo1
	/*	var x Sayer          // 声明一个Sayer类型的变量x
		a := Cat{}           // 实例化一个cat
		b := Dog{}           // 实例化一个dog
		x = a                // 可以把cat实例直接赋值给x
		fmt.Println(x.Say()) // 喵喵喵
		x = b                // 可以把dog实例直接赋值给x
		fmt.Println(x.Say()) // 汪汪汪*/

	// demo2
	/*	var x Mover
		// 使用值接收者实现接口之后，
		// 不管是dog结构体还是结构体指针 *dog类型的变量都可以赋值给该接口变量。
		// 因为Go语言中有对指针类型变量求值的语法糖，
		// dog 指针 fugui内部会自动求值 *fugui。
		var wangcai = Dog{} // 旺财是 dog类型
		x = wangcai         // x可以接收 dog类型
		x.move()
		var fugui = &Dog{} // 富贵是 *dog类型
		x = fugui          // x可以接收 *dog类型
		x.move()*/

	// demo3
	/*	var x Mover
		var wangcai = Dog{} // 旺财是dog类型
		x = wangcai         // x不可以接收dog类型
		// 此时实现Mover接口的是*dog类型，所以不能给x传入dog类型的wangcai，
		// 此时x只能存储*dog类型的值。
		var fugui = &Dog{} // 富贵是*dog类型
		x = fugui          // x可以接收*dog类型
		x.move2()*/

	// demo4
	/*	var peo People = Student{}
		think := "bitch"
		// 此处不能通过编译 因为 Speak方法的调用者应是*Student类型
		// 而 Student并没有实现接口 是 *Student类型实现接口
		fmt.Println(peo.Speak(think))
	*/

	// demo5
	// dog既可以实现Sayer接口，也可以实现Mover接口。
	/*	var x Sayer
		var y Mover
		var a = Dog{}
		x = a
		y = a
		fmt.Println(x.Say())
		y.move()*/

	// demo6
	/*	var x Mover
		var a = dog{name: "旺财"}
		var b = car{brand: "保时捷"}
		x = a
		x.move()
		x = b
		x.move()*/

	// demo7
	/*	var x WashingMachine
		var haier = haier{dryer: dryer{name: "xxxx甩干机"}}
		x = haier
		x.wash()
		x.dry()*/

	// demo8
	/*	var x animal
		x = cat{name: "花花"}
		x.move()
		x.say()*/

	// demo9
	// 定义一个空接口x 使用空接口实现可以接收任意类型的函数参数。
	/*	var x interface{}
		s := "pprof.cn"
		x = s
		fmt.Printf("type:%T value:%v\n", x, x)
		i := 100
		x = i
		fmt.Printf("type:%T value:%v\n", x, x)
		b := true
		x = b
		fmt.Printf("type:%T value:%v\n", x, x)*/

	// 空接口作为map值
	// map 键为字符串 值为接口
	// 使用空接口实现可以保存任意值的字典。
	/*	var studentInfo = make(map[string]interface{})
		studentInfo["name"] = "李白"
		studentInfo["age"] = 18
		studentInfo["married"] = false
		fmt.Println(studentInfo)*/

	var x interface{}
	x = 123
	// x.(T)
	// 该语法返回两个参数，第一个参数是x转化为T类型后的变量，
	// 第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。
	v, ok := x.(int)
	if ok {
		fmt.Println(v)
		fmt.Printf("%T", v)
	} else {
		fmt.Println("类型断言失败")
	}
	// 因为空接口可以存储任意类型值的特点，所以空接口在Go语言中的使用十分广泛。
	// 关于接口需要注意的是，只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。
	// 不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗。

}
