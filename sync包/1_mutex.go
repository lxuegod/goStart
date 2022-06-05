package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//互斥锁
	var lock sync.Mutex
	go func() {
		//加锁
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("func1 get lock at " + time.Now().String())
		time.Sleep(time.Second)
		fmt.Println("func1 release lock " + time.Now().String())
	}()

	time.Sleep(time.Second / 10)

	go func() {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("func2 get lock at " + time.Now().String())
		time.Sleep(time.Second)
		fmt.Println("func1 release lock " + time.Now().String())
	}()

	//等待所有goroutine执行完毕
	time.Sleep(time.Second * 4)
}
