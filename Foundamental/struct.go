package main

import (
	"encoding/json"
	"fmt"
)

// 类型别名
// 类型定义
type NewInt int

// 类型别名
type MyInt = int

type person struct {
	name string
	city string
	age  int8
}

type student struct {
	name string
	age  int
}

type Student struct {
	ID     int
	Gender string // json序列化是默认使用字段名作为 key
	Name   string // 私有不能被 json包访问
}

func (n NewInt) SayHello() {
	fmt.Println("hello 我是 int")
}

// go没有自带的构造函数，手写构造函数
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

// 结构体允许其成员字段在声明时没有字段名而只有类型，称为匿名字段
type person2 struct {
	string
	int
}

// 方法与函数最大区别在于，方法限定了使用方法的对象
// 方法 接受体可为指针类型也可也是非指针类型 指针类型的话，会在函数体中修改其值
func (p person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

// 方法 方法可以添加返回值 但实际调用时不一定需要接受返回值
func (p *person) SetAge(newAge int8) string {
	p.age = newAge
	return p.name
}

// Address 地址结构体
type Address struct {
	Province   string
	City       string
	CreateTime string
}

type Email struct {
	Account string
	// 字段名与 Address 冲突
	CreateTime string
}

// User 用户结构体
type User struct {
	Name   string
	Gender string
	// 一个结构体可以嵌套包含另一个机构提或结构体指针
	//Address Address
	Address // 匿名结构体
	Email
}

// Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

type Dog struct {
	Feet    int8
	*Animal // 嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪！\n", d.name)
}

// Student 学生
// 结构体字段值可以公开访问，小写表示私有
type student2 struct {
	ID     int    `json:"tag_ID"` // 通过 tag实现 json序列化该字段的key
	Gender string // json序列化是默认使用字段名作为key
	Name   string // 私有不能被 json包访问
}

// Class 班级
type Class struct {
	Title    string
	Students []*Student
}

func demo(ce []student) {
	//切片是引用传递，是可以改变值的
	ce[1].age = 999
	// ce = append(ce, student{3, "xiaowang", 56})
	// return ce
}

func main() {
	var a NewInt
	var b MyInt

	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int

	var p1 person
	fmt.Printf("p1=%T\n", p1)
	p1.name = "lzc"
	p1.city = "beijing"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)

	fmt.Printf("p1=%#v\n", p1)

	// 临时数据结构场景 可以使用匿名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "lzc"
	user.Age = 18
	fmt.Printf("user=%#v\n", user)

	// p2 是结构体指针
	var p2 = new(person)
	fmt.Println(p2)
	fmt.Printf("%T\n", p2)     //*main.person
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}

	// p3 是结构体指针
	var p3 = new(person)
	fmt.Println(&p3)
	p3.name = "测试"
	p3.age = 18
	p3.city = "北京"
	fmt.Printf("p3=%#v\n", p3) //p2=&main.person{name:"测试", city:"北京", age:18}

	// p4 对结构体进行取地址操作 相当于对结构体类型进行了一次 new实例化操作
	p4 := &person{}
	fmt.Println(&p4)
	fmt.Printf("%T\n", p4)     //*main.person
	fmt.Printf("p4=%#v\n", p4) //p3=&main.person{name:"", city:"", age:0}
	p4.name = "博客"
	p4.age = 30
	p4.city = "成都"
	fmt.Printf("p4=%#v\n", p4) //p3=&main.person{name:"博客", city:"成都", age:30}

	var p5 person
	fmt.Printf("p5=%#v\n", p5)

	p6 := person{
		name: "lzc",
		city: "wh",
		age:  18,
	}
	fmt.Printf("p6=%#v\n", p6)
	fmt.Printf("p6=%v\n", p6)

	// 对结构体指针初始化
	p7 := &person{
		name: "lzc",
		city: "wh",
		age:  18,
	}
	fmt.Printf("p7=%#v\n", p7)

	// 某些字段可以不指定初始值
	p8 := &person{
		name: "lzc",
	}
	fmt.Printf("p8=%#v\n", p8)

	p9 := &person{
		"xxxx",
		"bj",
		18,
	}
	fmt.Printf("p9=%#v\n", p9)

	type test struct {
		a int8
		b int8
		c int8
	}

	n := test{
		1, 2, 3,
	}
	// 加上#还会打印类型
	fmt.Printf("n=%#v\n", &n.a)
	fmt.Printf("n=%#v\n", &n.b)
	fmt.Printf("n=%#v\n", &n.c)

	// map 存储 字符串类型的键 结构体指针值
	m := make(map[string]*student)
	// 初始化一个 结构体切片
	stus := []student{
		{
			name: "lzc",
			age:  18,
		},
		{
			name: "lzc2",
			age:  18,
		},
		{
			name: "lzc3",
			age:  18,
		},
	}
	// 遍历切片
	for _, stu := range stus {
		//
		m[stu.name] = &stu
	}
	// v 是结构体指针
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
		fmt.Println(k, "=>", v.age)
	}

	p10 := newPerson("lzc", "cs", 24)
	fmt.Printf("p10=%#v\n", p10)

	p10.Dream()

	var p11 person
	fmt.Printf("%T\n", p11)
	p11.name = "lzc"

	p11.Dream()
	pName := p11.SetAge(50)
	fmt.Println(p11.age)
	fmt.Println(pName)

	var p12 *person
	p12 = &p11
	p12.Dream()

	p12.SetAge(100)
	fmt.Println(p11.age)
	fmt.Println(p12.age)

	var n1 NewInt
	n1.SayHello()
	n1 = 100
	fmt.Printf("n1=%#v, %T\n", n1, n1)

	p13 := person2{
		"lzc",
		18,
	}
	fmt.Printf("p13=%#v\n", p13)
	fmt.Println(p13.string, p13.int)

	user1 := User{
		Name:   "pprof",
		Gender: "女",
		Address: Address{
			Province: "黑龙江",
			City:     "哈尔滨",
		},
	}
	fmt.Printf("user1=%#v\n", user1)

	var user2 User
	user2.Name = "lzc"
	user2.Gender = "男"
	user2.Province = "北京"
	// 直接访问匿名结构体字段名
	user2.City = "北京"
	// 直接通过匿名结构体.字段名访问
	user2.Address.City = "武汉"
	fmt.Printf("user2=%#v\n", user2)

	var user3 User
	user3.Name = "lzc"
	user3.Gender = "女"

	user3.Address.CreateTime = "1000"
	user3.Email.CreateTime = "2000"
	fmt.Printf("user3=%#v\n", user3)

	var a1 Animal
	a1.name = "lzc"
	a1.move()

	var d1 Dog
	d1 = Dog{
		Feet: 4,
		Animal: &Animal{
			name: "lzc2",
		},
	}
	d2 := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "lzc2",
		},
	}
	d1.wang()
	d1.move()
	d2.wang()
	d2.move()

	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}

	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}

	fmt.Println(c)
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	fmt.Printf("json:%T\n", data)

	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	fmt.Printf("%T\n", str)
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)

	s1 := student2{
		ID:     1,
		Gender: "男",
		Name:   "lzc",
	}
	data, err = json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json:%s\n", data)

	var ce []student // 定义切片
	ce = []student{
		student{"lzc", 18},
		student{"lzc2", 22},
	}
	fmt.Println(ce)
	demo(ce)
	fmt.Println(ce)

	ce2 := make(map[int]student)
	ce2[1] = student{"lzc", 18}
	ce2[2] = student{"lzc2", 18}
	fmt.Println(ce2)
	delete(ce2, 2)
	fmt.Println(ce2)
}
