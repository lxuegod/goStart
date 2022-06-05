package main

import "fmt"

//实现两数相加
//面向过程
func Add1(a, b int) int {
	return a + b
}

//自定义类型
type long int

func (tmp long) Add2(other long) long {
	return tmp + other
}

func main() {
	var result int
	result = Add1(1, 1) //普通函数调用方式
	fmt.Println("result = ", result)

	//定义一个变量
	var a long = 2
	//调用方法格式：变量名.函数（所需参数）
	r := a.Add2(3)
	fmt.Println("r = ", r)
}
