package main

import (
	"fmt"
	"time"
)

//全局变量创建一个channel
var ch = make(chan int)

//定义一个打印机 参数为字符串，按每个字符打印
//打印机属于公共资源
func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func Person1() {
	Printer("hello")
	ch <- 666 //给管道写数据，发送
}

func Person2() {
	<-ch //从管道取数据，接收，如果管道没有数据它就会阻塞
	Printer("world")
}

func main() {
	//新建2个协程，代表2个人，2个人同时使用打印机
	go Person1()
	go Person2()

	//特地不让主协程结束，死循环
	for {

	}
}
