package main

import (
	"fmt"
	"time"
)

//ch只写，quit只读
func fibonacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		//监听channel数据流动的方向
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Printf("\nflag = %t", flag)
			return
		}
	}
}
func Main1() {
	ch := make(chan int)    //数字通信
	quit := make(chan bool) //程序是否结束

	//消费者，从channel读取内容
	//新建协程
	go func() {
		for i := 0; i < 8; i++ {
			num := <-ch
			fmt.Printf("%5d", num)
		}
		//可以停止
		quit <- true
	}()
	//生产者生产数字，写入channel
	fibonacci(ch, quit)
}

//select的超时处理
func Main2() {
	ch := make(chan int)
	quit := make(chan bool)

	//新开一个协程
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second):
				fmt.Println("超时")
				quit <- true
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	<-quit
	fmt.Println("程序异常结束")
}

func main() {
	Main2()
}
