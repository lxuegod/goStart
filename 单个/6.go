package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte
	addr string
	age  int
}

func main() {
	//指针有合法指向性后 才能操作成员函数
	//先定义一个普通变量的结构体
	var s Student
	//再定义一个指针变量  保存s的地址
	var p1 *Student
	p1 = &s
	//通过指针操作成员  p1.id 和 (*p1).id完全等价 没有->运算符 只有.运算符
	p1.id = 18
	(*p1).name = "mike"
	p1.sex = 'm'
	p1.age = 18
	p1.addr = "bj"
	fmt.Println("p1 = ", p1)

	//通过new申请一个结构体
	p2 := new(Student)
	p2.id = 18
	(*p2).name = "mike"
	p2.sex = 'm'
	p2.age = 18
	p2.addr = "bj"
	fmt.Println("p2 = ", *p2) //加*后  打印的&不见了

}
