package main

import "fmt"

type Humaner interface {
	//方法，只有声明，没有实现，由别的类型（自定义）实现
	sayhi()
}

//超集  里面包含其他的接口Humaner  Humaner就为子集
//超集可以转化为子集
type Student struct {
	name string
	id   int
}

//Student的方法
func (tmp *Student) sayhi() {
	fmt.Printf("Student [%s,%d] sayhi\n", tmp.name, tmp.id)
}

type Teacher struct {
	addr  string
	group string
}

//Teacher的方法
func (tmp *Teacher) sayhi() {
	fmt.Printf("Teacher [%s,%s] sayhi\n", tmp.addr, tmp.group)
}

//自定义的类型
type Mystr string

//Mystr实现了此方法
func (tmp *Mystr) sayhi() {
	fmt.Printf("Mystr [%s] sayhi", *tmp)
}

//定义一个普通函数，函数的参数为接口类型
//只有一个函数，可以有不同表现，多态
func WhoSayHi(i Humaner) {
	i.sayhi()
}

func main() {
	s := &Student{"mike", 1234}
	t := &Teacher{"bj", "go"}
	var str Mystr = "hello mike"

	//调用同一函数，不同表现，多态，多种形态
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)

	//创建一个切片
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str

	//切片的下标 第一个返回下标，第二个返回下标对应的值
	for _, i := range x {
		//j := 0
		fmt.Printf("\n")
		i.sayhi()
		//j++
	}
}

func main01() {
	//定义接口类型的变量
	var i Humaner

	//只要实现了此接口方法的类型，那么这个类型的变量（接收者类型）就可以赋值给i
	s := &Student{"mike", 1234}
	i = s
	i.sayhi()

	t := &Teacher{"bj", "go"}
	i = t
	i.sayhi()

	var str Mystr = "hello mike"
	i = &str
	i.sayhi()
}
