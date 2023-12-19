package main

import "fmt"

type Int int

func (i Int) SayHello() {
	fmt.Println("hello, i am ", i)
}

func main() {
	var num int = 10
	fmt.Printf("num: %v\n", num)
	var num2 Int
	num2.SayHello()

	b := "aaaa"
}
