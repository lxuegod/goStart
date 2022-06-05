package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (p Person) SetInfoValue() {
	fmt.Println("SetInfoValue :%p, %v\n", &p, p)
}

func (p *Person) SetInfoPointer() {
	fmt.Println("SetInfoPointer :%p, %v\n", &p, p)
}

func main() {
	p := Person{"mike", 'm', 18}
	fmt.Printf("main :%p, %v", &p, p)

	// p.SetInfoPointer() //传统调用方式
	// //保存方式入口地址
	// pFunc := p.SetInfoPointer()
	// //这个就是方法值，调用函数时，无须再传递接受者，隐藏了接收者

	// pFunc() //等价于p.SetInfoPointer()

	//方法表达式
	f := (*Person).SetInfoPointer
	f(&p) //显式把接收者传递过去  p.SetInfoPointer()

	f2 := (Person).SetInfoValue
	f2(p) //显式把接收者传递过去  p.SetInfoValue()
}
