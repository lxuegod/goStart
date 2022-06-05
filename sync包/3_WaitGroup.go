package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	//添加等待goroutine数量为5
	waitGroup.Add(5)

	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println("work " + strconv.Itoa(i) + " is done at " + time.Now().String())
			//等待1s后减少等待数1
			time.Sleep(time.Second)
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()
	fmt.Println("all works are done at " + time.Now().String())
}
