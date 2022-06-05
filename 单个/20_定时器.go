package main

import (
	"fmt"
	"time"
)

func main() {
	Main1()
}

//time.NewTimer()时间到了只会响应一次
func Main1() {
	//创建一个定时器，设置时间为2s，2s后往通道里写内容（当前时间）
	now := time.Now()
	formatNow := now.Format("2006-01-02 15:04:05")
	timer := time.NewTimer(2 * time.Second)
	fmt.Printf("当前时间：%v\n", formatNow)

	//2s后往timer.C中写数据，有数据后就能读取
	t := <-timer.C //channel没有数据前阻塞
	fmt.Println("t = ", t)
}

//延时的三种方式

func Main2() {
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("时间到")
}

func Main3() {
	time.Sleep(2 * time.Second)
	fmt.Println("时间到")
}

func Main4() {
	<-time.After(2 * time.Second) //定时2s,阻塞2s,2s后产生一个事件,往channel写内容
	fmt.Println("时间到")
}

//停止定时器
func Main5() {
	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("子线程可以打印了，因为定时器的时间到")
	}()
	timer.Stop()
	for {

	}
}

//重置定时器
func Main6() {
	timer := time.NewTimer(3 * time.Second)
	ok := timer.Reset(1 * time.Second) //重新设置为1s
	fmt.Println("0k = ", ok)
	<-timer.C
	fmt.Println("时间到")
}

//ticker的使用
//ticker只要定义完成，从此刻开始计时，不需要任何其他的操作，每隔固定时间都会触发
//使用timer定时器，超时后需要重置，才能继续触发
func Main7() {
	ticker := time.NewTicker(1 * time.Second)

	i := 0
	for {
		<-ticker.C
		i++
		fmt.Println("i = ", i)
		if i == 5 {
			ticker.Stop()
			break
		}
	}
}
