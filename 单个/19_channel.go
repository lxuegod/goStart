package main

import (
	"fmt"
	"time"
)

func Main1() {
	//创建channel
	ch := make(chan string)

	defer fmt.Println("主线程也结束")

	go func() {
		defer fmt.Println("子线程也调用完毕")
		for i := 0; i < 2; i++ {
			fmt.Println("子线程 i = ", i)
			time.Sleep(time.Second)
		}
		ch <- "我是子线程，工作完毕"
	}()
	str := <-ch //没有数据前，阻塞
	fmt.Println("str = ", str)
	fmt.Println("Main1()结束了")
}

//无缓存的channel
func Main2() {
	//创建一个无缓存的channel
	ch := make(chan int, 0)

	//len(ch) 缓存区剩余数据个数  cap(ch)缓存区大小
	fmt.Printf("len(ch) =%d ,cap(ch) =%d\n ", len(ch), cap(ch))

	//新建协程
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("子协程i =%d\n ", i)
			ch <- i //往chanl中写内容
			fmt.Printf("len(ch) =%d ,cap(ch) =%d\n ", len(ch), cap(ch))
		}
	}()
	//延时
	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-ch //读管道的内容，没有内容阻塞
		fmt.Printf("num = %d\n", num)
	}
	fmt.Println("Main2()结束了")
}

//创建一个有缓存的channel
func Main3() {
	//创建一个有缓存的channel
	ch := make(chan int, 3)

	//len(ch) 缓存区剩余数据个数  cap(ch)缓存区大小
	fmt.Printf("len(ch) =%d ,cap(ch) =%d\n ", len(ch), cap(ch))

	//新建协程
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i //往chanl中写内容
			fmt.Printf("子协程[%d]:len(ch) =%d ,cap(ch) =%d\n ", i, len(ch), cap(ch))
		}
	}()
	//延时
	time.Sleep(5 * time.Second)

	for i := 0; i < 10; i++ {
		num := <-ch //读管道的内容，没有内容阻塞
		fmt.Printf("num = %d\n", num)
	}
	fmt.Println("Main3()结束了")
}

//关闭channel
func Main4() {
	ch := make(chan int) //创建一个无缓存的channel

	//新建一个goroutine
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //往通道写数据
		}
		close(ch) //不需要写数据时关闭channel
	}()

	// //第一种方法
	// for {
	// 	if num, ok := <-ch; ok == true {
	// 		fmt.Println("num = ", num)
	// 	} else { //管道关闭
	// 		break
	// 	}
	// }
	//第二种用range遍历
	for num := range ch {
		fmt.Println("num = ", num)
	}
}

// func Main5() {
// 	//创建channel默认是双向的
// 	ch := make(chan int)

// 	//双向channel能隐式转换为单向的  定义方式
// 	var writeCh chan<- int = ch //只能写，不能读
// 	var readCh <-chan int = ch  //只能读，不能写

// 	//只能由双向转化为单向  不能由单向转化为双向

// }

//单向channel的应用

//此通道只能写，不能读
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

//此通道只能读，不能写
func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("num = ", num)
	}
}

func Main5() {
	//创建一个双向通道
	ch := make(chan int)

	//生产者，生产数字，写入通道
	//新开一个协程
	go producer(ch) //channel传参，引用传递

	//消费者  从channel读取内容，打印
	consumer(ch)
}

func main() {
	Main5()
}
