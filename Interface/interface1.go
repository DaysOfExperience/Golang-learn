package main

import "fmt"

// 接口/Interface: 接口（interface）定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节。
// 其实就是, 一个接口类型规定了某些方法, 如果一个类实现了这些方法, 那么它就是一个接口类型的对象
// 比如我定义学生接口: 1.学习 2.睡觉 只要某个struct实现了这两个方法, 那它就符合学生接口

// 相较于之前章节中讲到的那些具体类型（字符串、切片、结构体等）更注重“我是谁”，接口类型更注重“我能做什么”的问题。
// 接口类型就像是一种约定——概括了一种类型应该具备哪些方法，在Go语言中提倡使用面向接口的编程方式实现解耦。
// 一个接口类型就是一组方法的集合，它规定了需要实现的所有方法。

// 每个接口类型由任意个方法签名组成
type Person interface {
	eat()   // 一个人, 需要会吃
	sleep() // 一个人, 需要会睡
}

// 学生会吃会睡会学习, 前两个条件满足, 它就实现了Person接口
type Student struct {
	name string
	age  int
}

func (s *Student) eat() {
	fmt.Println(s.name, "ate")
}
func (s *Student) sleep() {
	fmt.Println(s.name, "slept")
}
func (s *Student) study() {
	fmt.Println(s.name, "studying")
}

// 工人会吃会睡会工作, 前两个条件满足, 它就实现了Person接口
type Worker struct {
	name string
	age  int
}

func (w *Worker) eat() {
	fmt.Println(w.name, "ate")
}
func (w *Worker) sleep() {
	fmt.Println(w.name, "slept")
}
func (w *Worker) work() {
	fmt.Println(w.name, "working")
}

// 实现一个方法:饿肚子并且瞌睡, 则函数内需要调用对应类型的吃 && 睡, 那么很多类型其实都会吃 && 睡
// 但是我们可不能针对每一个会吃/睡的类型都实现一个函数, 此时将参数实现为接口类型即可
// 这样所有实现了eat和sleep方法的类都可以传给这个参数
func HungryAndSleepy(p Person) {
	p.eat()
	p.sleep()
}
func test1() {
	w := &Worker{
		"一个程序猿",
		22,
	}
	s := &Student{
		"苦逼学生",
		18,
	}
	HungryAndSleepy(w)
	HungryAndSleepy(s)
}

// 接口类型区别于那些具体类型，让我们专注于该类型提供的方法，而不是类型本身。

// 再举个例子: 比如支付宝和微信类都实现了支付的方法, 那么有一个超市结算的函数, 需要一个参数, 也就是支付宝或者微信对象
// 内部调用其支付的方法
// 这就是典型的“不关心它是什么，只关心它能做什么”的场景。  并不关心这是哪个支付系统, 只要它能支付即可~
// 比如一个网上商城可能使用支付宝、微信、银联等方式去在线支付，我们能不能把它们当成“支付方式”来处理呢？
// 比如三角形，四边形，圆形都能计算周长和面积，我们能不能把它们当成“图形”来处理呢？
// 比如满减券、立减券、打折券都属于电商场景下常见的优惠方式，我们能不能把它们当成“优惠券”来处理呢？
// 这其实就是抽象~~~~

// 一个接口类型的变量能够存储所有实现了该接口的类型的对象。
// 值接收者和指针接收者
// 1. 类型使用值接收者实现接口所规定的方法之后，不管是结构体类型对象还是对应的结构体指针类型的变量都可以赋值给该接口变量。
// 2. 类型使用指针接收者实现接口所规定的方法之后, 只能将结构体指针类型的变量赋值给该接口的变量
type Animal interface {
	say()
}
type dog struct {
}

func (d *dog) say() {
	fmt.Println("汪汪汪")
}

type cat struct {
}

func (c *cat) say() {
	fmt.Println("喵喵喵")
}
func test2() {
	var animal Animal // 接口类型变量的定义
	animal = &dog{}   // 这里必须是结构体指针, 如果上方实现的say不是指针类型接收, 则结构体对象可以
	// 也就是说, 这里的d和上面的say接收的d *dog需要对应匹配
	animal.say()
	animal = &cat{}
	animal.say()
}

// 类型和接口的关系
// 多种类型可以实现同一接口
// 一个类型可以实现多个接口
// 接口两个方法, 类型实现了方法2, 类型的数据成员是另一个类型, 则若另一个类型实现了方法1, 则视为类型实现了该接口
type inter interface {
	fun1()
	fun2()
}
type s1 struct{}

func (s s1) fun1() {}

// s2类型实现了inter接口
type s2 struct {
	s1
}

func (s s2) fun2() {}
func test3() {
	var i inter
	i = s2{
		s1: s1{},
	}
	//i = s1{ // Cannot use 's1{ }' (type s1) as the type inter Type does not implement 'inter' as some methods are missing: fun2()
	//	// s1没有实现fun2方法, 没有实现inter接口
	//}
	i.fun2()
}

// 接口组合
// 接口与接口之间可以通过互相嵌套形成新的接口类型
// 结构体也可以将接口作为自己的一个字段

// 空接口
// 空接口是指没有定义任何方法的接口类型。因此任何类型都可以视为实现了空接口。
// 也正是因为空接口类型的这个特性，空接口类型的变量可以存储任意类型的值。
func test4() {
	var x interface{} // 空接口类型的变量, 存储任意类型的值
	x = "你好"          // 字符串型
	fmt.Printf("type:%T value:%v\n", x, x)
	x = 100 // int型
	fmt.Printf("type:%T value:%v\n", x, x)
	x = true // 布尔型
	fmt.Printf("type:%T value:%v\n", x, x)
	x = Student{"111", 2}                  // 结构体类型
	fmt.Printf("type:%T value:%v\n", x, x) // main.Student {111 2}
}

// 空接口的应用: 空接口类型的变量可以存储任意类型的值, 其实很牛逼的
// 1.空接口作为函数参数, 则可以接收任意类型的实参
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

// 2. 空接口类型作为map的值
func test5() {
	m := map[string]interface{}{}
	m["int"] = 1
	m["string"] = "sssstring"
	m["map"] = map[int]int{1: 2, 2: 4}
	m["bool"] = false
	fmt.Println(m)
}

// 接口值
// 由于接口类型的值可以是任意一个实现了该接口的类型值，所以接口值除了需要记录具体值之外，还需要记录这个值属于的类型。
// 也就是说接口值由“类型”和“值”组成，鉴于这两部分会根据存入值的不同而发生变化，我们称之为接口的 动态类型 和 动态值。

// 接口变量的零值为 类型为nil,值为nil
// 接口值是支持相互比较的，当且仅当接口值的动态类型和动态值都相等时才相等。
func test6() {
	// Person Student Worker
	var p Person          // 接口的零值 : nil nil
	fmt.Println(p == nil) // true   // 当接口变量的动态类型和动态值都为nil时才为true
	//p.sleep()  // panic
	p = &Student{"111", 11} // 动态类型动态值都不为nil
	var w *Worker           // nil
	p = w                   // 动态类型不为nil, 动态值为nil
	fmt.Println(p == nil)   // false
	// 接口值是支持相互比较的，当且仅当接口值的动态类型和动态值都相等时才相等。
}

// 类型断言
// 接口值可能赋值为任意类型的值，那我们如何从接口值获取其存储的具体数据呢？
func test7() {
	var p Person // 这是一个接口变量
	p = &Worker{
		"111",
		11,
	}
	fmt.Printf("%T\n", p) // p的类型打印出来
	// fmt包内部使用反射的机制在程序运行时获取到动态类型的名称。 反射之后再说

	//类型断言:
	value, ok := p.(*Student) // 若断言失败, 则value就是对应类型的零值
	if ok {
		fmt.Println("p的动态类型是*Student类型, vlaue: ", value)
	} else {
		fmt.Println("p的动态类型不是*Student类型, value: ", value)
	}
}

// 只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。
// 切记不要为了使用接口类型而增加不必要的抽象，导致不必要的运行时损耗。
// 接口是一种类型，一种抽象的类型。
// 摘自gin框架routergroup.go
type IRouter interface {
	//...
}

type RouterGroup struct {
	//...
}

var _ IRouter = &RouterGroup{} // 确保RouterGroup实现了接口IRouter
// 如果 RouterGroup 没有实现 IRouter 中定义的所有方法，程序将无法通过编译。
func main() {
	//test1()
	//test2()
	//test4()
	//test5()
	test7()
}
