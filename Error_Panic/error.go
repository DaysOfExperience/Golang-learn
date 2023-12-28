package main

import (
	"errors"
	"fmt"
	"os"
)

//	type error interface {
//		Error() string
//	}
func test1() {
	_, err := os.Open("./xx.go")
	// err是一个接口类型的变量, 默认值为nil, 若返回的err不是nil, 说明发生了错误
	if err != nil {
		fmt.Printf("open file failed, err msg:%v\n", err.Error())
	}
}

func test2() {
	// 创建一个错误
	// func New(text string) error
	err := errors.New("test error")
	fmt.Println(err.Error())

	n, err := fun(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
	}
	n, err = fun(-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
	}

	// 2.  fmt.Errorf
	e := fmt.Errorf("test error second, %d", 2)
	fmt.Println(e.Error())
}
func fun(num int) (int, error) {
	// 此方法, 用于返回正整数的相反值
	if num <= 0 {
		return 0, errors.New("Invalid num")
	}
	return -num, nil
}

// 此结构体实现了error接口
type myError struct {
	msg string
}

func (e *myError) Error() string {
	return e.msg
}

type MyError struct {
	code int
	msg  string
}

func (m MyError) Error() string {
	return fmt.Sprintf("code:%d, msg:%v", m.code, m.msg)
}
func NewError(code int, msg string) error {
	return MyError{
		code: code,
		msg:  msg,
	}
}

// MyError实现了error, 就应该这样做, 而不是实现为MyError的方法
func Code(err error) int {
	if e, ok := err.(MyError); ok {
		return e.code
	}
	return -1
}
func Msg(err error) string {
	if e, ok := err.(MyError); ok {
		return e.msg
	}
	return err.Error()
}
func test3() {
	err := NewError(100, "test MyError")
	fmt.Println(err.Error())
	fmt.Println(Code(err), Msg(err))
}
func main() {
	test1()
	test2()
	test3()
}
