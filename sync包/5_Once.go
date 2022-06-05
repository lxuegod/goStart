package main

import (
	"fmt"

	"sync"
)

var once sync.Once
var waitGroup sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			once.Do(OnlyOnce)
		}()
	}
	waitGroup.Wait()
}

func OnlyOnce() {
	fmt.Println("only once")
}
