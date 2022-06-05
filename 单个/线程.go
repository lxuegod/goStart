//主线程退出了，子线程也跟着退出
//用go关键字新建子协程

//runtime 包

package main

import (
	"fmt"
	"runtime"
)

func main() {
	func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go")
		}
	}()
	for i := 0; i < 2; i++ {
		//Gosched让出时间片，先让别的协程执行，它执行完，再执行此协程

		runtime.Gosched()
		fmt.Println("hello")

		//终止协程runtime.Goexit()
		//指定核GOMAXPROCS
		n := runtime.GOMAXPROCS(4)
		fmt.Println("n = ", n)
	}
}
