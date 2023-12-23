package main

import (
	"fmt"
	"reflect"
)

func test1(x interface{}) {
	// x是一个接口变量, 函数执行时并不知道其类型和值, 利用反射可以获取
	// 1. 类型断言, 太麻烦
	// 2. 利用反射
	obj := reflect.TypeOf(x) // 返回值类型为Type接口类型, 他有很多方法
	fmt.Println("obj:", obj, "obj.Name():", obj.Name(), "obj.Kind:", obj.Kind())
	//fmt.Printf("Type of obj: %T\n", obj) // 把这个obj的类型输出出来  *reflect.rtype

	// reflect.ValueOf()返回的是reflect.Value类型，其中包含了原始值的值信息。reflect.Value与原始值之间可以互相转换。
	value := reflect.ValueOf(x) // 返回一个Value类型的结构体对象
	fmt.Printf("%T %v\n", value, value)
	kind := value.Kind()
	// 通过reflect.Value获取原始值
	switch kind {
	case reflect.Int:
		ret := int(value.Int())
		fmt.Printf("%T %v\n", ret, ret)
	case reflect.Float32:
		ret := float32(value.Float())
		fmt.Printf("%T %v\n", ret, ret)
	case reflect.Float64:
		ret := float64(value.Float()) // redundant conversion
		fmt.Printf("%T %v\n", ret, ret)
	}
}
func test3(x interface{}) {
	// 可以获取接口变量的type kind value, 还有原始类型的值
	// 如何修改?
	// 如果传值, 则会panic报错
	// 必须传指针
	v := reflect.ValueOf(x) // 获取Value
	k := v.Elem().Kind()    // 获取类型
	switch k {
	case reflect.Int:
		v.Elem().SetInt(100)
	}
}

type Student struct{}
type Worker struct{}

func test2(i int64) {
	fmt.Println(i)
}

// reflect.TypeOf和reflect.ValueOf两个函数来获取任意对象的Value和Type
// Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。
func main() {
	//test1(1)
	//test1(float32(2.1))
	//test1(2.2)
	//test1(int64(11))
	//test1(true)
	//test1("ssss")
	//test1(map[int]int{1: 2})
	//test1(Student{})
	//test1(Worker{})
	//test1(&Student{})
	//test1(&Worker{})

	//test2(int32(1))

	var x int = 10
	test3(&x)
	fmt.Println(x)
}
