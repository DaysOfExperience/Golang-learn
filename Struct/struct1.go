package main

import "fmt"

// Go语言没有类的概念, 只有结构体, 也没有C++的继承的概念

// Go语言中可以使用type关键字来定义自定义类型。自定义类型是定义了一个全新的类型。
// 类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型。
// 类型别名只会在代码中存在，编译完成时并不会有这个类型。  可以理解为某些类型太长我们可以设定一个类型别名, 只是为了可读性以及便于编程
// 而type定义类型, 就是一个新的类型, 即使它是基于int定义的

// 结构体
type person struct {
	// name, city string   // 同类型可以写在一行里面
	name string
	age  int
	city string
}

func test1() {
	// 结构体的实例化
	var p person
	p.name = "yyy"
	p.age = 19
	p.city = "China"

	// 匿名结构体
	// var user struct{num1 int; num2 int}

	// 使用new创建结构体实例, 不好用
	var p2 = new(person)            // 返回的是指针
	p2.name = "zzz"                 // 结构体的指针可以直接用.访问成员, 语法糖
	fmt.Printf("%T, %#v\n", p2, p2) // *main.person &main.person{name:"zzz", age:0, city:""}

	var p3 = &person{ // 结构体指针
		name: "zzz", // 键值对初始化
		age:  10,
		city: "hhh",
	}
	// 省略键值对的键, 使用值的列表进行初始化, 此时必须全部初始化, 且顺序需要对应, 不能混用键值对初始化
	var p4 = person{ // 结构体对象
		"zzz",
		11,
		"sss",
	}
	fmt.Printf("%T, %#v\n", p3, p3)
	fmt.Printf("%T, %#v\n", p4, p4)
}

// 结构体的内存布局  略

// 结构体初始化之 - 构造函数
// Golang没有像C++一样的构造函数, 所以我们只能创建一个函数, 用于构造一个结构体实例化对象
func newPerson(name string, age int, city string) *person {
	return &person{
		name,
		age,
		city,
	}
}

// Golang中的方法&接收者
// Go没有C++那样的成员构造函数, 成员方法怎么定义呢? 都是定义在类外部的, 可以指定这个方法的接收者变量 接收者类型
// 这样这个方法只能对应类型的结构体变量调用!!!!

// 重点: 因为Golang中的结构体是值类型, 所以接收者类型分为两类: 1. 值类型的结构体变量 2. 结构体指针
// 显然, 指针传递更快, 且可以用于修改实例的属性
// func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
//     函数体
// }
// 差别只是函数名前指定一下这是哪个类的方法
func (p *person) changeName(s string) {
	p.name = s
}
func (p person) sayHello() {
	fmt.Printf("hello~\n")
}
func test2() {
	// 调用构造函数
	p := newPerson("zzz", 1, "zzzz") // *main.person, &main.person{name:"zzz", age:1, city:"zzzz"}
	fmt.Printf("%T, %#v\n", p, p)
}

/*
什么时候应该使用指针类型接收者
1. 需要修改接收者中的值
2. 接收者是拷贝代价比较大的大对象
3. 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
*/

// 任意类型添加方法
// 注意事项： 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。  所以内置的int我们不能给他定义方法
type newInt int

func (i newInt) say() {
	fmt.Println("fuck you")
}

// func main() {
// 	// test1()
// 	test2()
// }
