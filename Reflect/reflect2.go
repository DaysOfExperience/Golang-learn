package main

import (
	"fmt"
	"reflect"
)

// 结构体反射, 其实就是用接口变量接收了一个结构体对象, 此时如何获取其各种属性, 调用其方法

type Person struct {
	name  string `json:"name" ini:"s_name"`
	age   int    `json:"age" ini:"s_age"`
	score int    `json:"score" ini:"s_score"`
}

func (p *Person) Setage(age int) {
	p.age = age
}
func (p *Person) Getage() int {
	fmt.Println(p.age)
	return p.age
}
func test(x interface{}) {
	// x是一个接口变量, 此时传过来一个结构体对象
	//t := reflect.TypeOf(x)
	//fmt.Printf("name:%v kind:%v\n", t.Name(), t.Kind()) // name:Person kind: struct
	//for i := 0; i < t.NumField(); i++ {
	//	f := t.Field(i) //  StructField
	//	fmt.Println(f.Name)
	//	fmt.Println(f.Type)
	//	fmt.Println(f.Tag)
	//	fmt.Println("json tag:", f.Tag.Get("json"))
	//}

	// 除了通过索引获取, 还可以通过字段名获取
	t := reflect.TypeOf(x)
	filedObj, ok := t.FieldByName("name")
	if ok {
		fmt.Println("name field exists, type of name:", filedObj.Type) // 字段类型
	} else {
		fmt.Println("name field is not exists")
	}
}

// 上面是通过索引和字段名获取结构体对象类型的接口变量的各种属性字段
// 还可以调用其方法
func test4(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Println(t.NumMethod()) // 结构体的方法数量

	v := reflect.ValueOf(x)
	// 下面需要拿到方法, 所以需要用值信息
	for i := v.NumMethod() - 1; i >= 0; i-- {
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", v.Method(i).Type())
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		if t.Method(i).Name == "Setage" {
			args := []reflect.Value{reflect.ValueOf(20)}
			v.Method(i).Call(args)
		} else {
			args := []reflect.Value{}
			age := v.Method(i).Call(args)
			fmt.Printf("test4: return age:%v\n", age[0].Int())
		}
	}
}
func main() {
	p := Person{
		"111", 18, 100,
	}
	//test(p)
	test4(&p)
}
