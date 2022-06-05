package main

import (
	"fmt"
)

func main() {
	s := []byte{'1'}
	s = append(s, '0'+byte(0))
	fmt.Println("s = ", string(s))
}
