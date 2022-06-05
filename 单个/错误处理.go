package main

import (
	"errors"
	"fmt"
)

func MyDeiv(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		//错误接口的使用
		err = errors.New("分母不能为0")
		//panic("aaaa")  panic函数为异常处理的函数 可以打印  程序崩溃的时候会自动调用
	} else {
		result = a / b
	}
	return
}

//recover的使用
func testa() {
	fmt.Println("aaaaaaa")
}
func testb(x int) {
	//设置recover
	defer func() {
		//recover()  可以打印panic的错误信息
		//fmt.Println(recover())
		if err := recover(); err != nil { //产生了panic异常
			fmt.Println(err)
		}
	}() //()为调用匿名函数
	var a [10]int
	a[x] = 111 //当x为20时，导致数组越界，产生一个panic，程序崩溃
}

func testc() {
	fmt.Println("ccccccc")
}

func main() {
	result, err := MyDeiv(10, 0)
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println("result = ", result)
	}

	//recover的测试
	testa()
	testb(20)
	testc()
}
