package main

import (
	myPackage "/../Struct/Test/myPackage"
	"fmt"
)

// 结构体匿名字段
// 默认会采用类型名作为字段名. 故, 同一类型的字段只能有一个
// 像下面这样再指定一个name也可以
type unname struct {
	string
	int
	name string
}

// 嵌套结构体
type Address struct {
	Ciy      string
	Province string
}
type Person struct {
	Name    string
	Gender  string
	Address Address // 这种的定义为匿名字段就挺合适, 因为同名嘛, 而且也只会有一个
}

func test() {
	o1 := unname{ // 匿名结构体实例化
		"111",
		1,
		"zzz",
	}
	user1 := &Person{
		Name:   "yyy",
		Gender: "man",
		Address: Address{
			"111",
			"222",
		},
	}
	fmt.Println(o1, user1)
}

type t1 struct {
	num  int
	time string
}
type t2 struct {
	number int
	time   string
}
type t3 struct {
	name string
	t1
	t2
}

func test3() {
	o := t3{
		name: "hhh",
		t1: t1{
			1,
			"19:38",
		},
		t2: t2{
			number: 2,
			time:   "19:39",
		},
	}
	// num := o.time   // ambiguous
	time1 := o.t1.time
	time2 := o.t2.time
	fmt.Println(time1, time2)
}

// 结构体的继承

type Animal struct {
	name string
}

func (a *Animal) eat(food string) {
	fmt.Println(a.name, "eats", food)
}

// 要使 Wolf 能够调用 Animal 的 eat 方法，您需要将 Animal 结构体作为匿名字段嵌入到 Wolf 结构体中。
// 这样，Wolf 就可以直接访问 Animal 的方法了。
type Wolf struct {
	age int
	Animal
}

func (w Wolf) wolf() {
	// 狼会wolf
	fmt.Println("wolf~~~~~~")
}
func testz() {
	w := Wolf{
		1,
		Animal{
			"nnnname",
		},
	}
	w.eat("zzz")
}

// 探索字段可见性
func test4() {
	myPackage.sayHello()
	myPackage.SayHello()
	num := myPackage.number
	num2 := myPackage.Number
	fmt.Println(num, num2)
	test2()
}

// Tag
func main() {
	// test()
	// test1()
	// test2()
	//test3()
	testz()
}
