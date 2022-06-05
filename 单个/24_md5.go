package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	str := "abc"
	x := md5.New()
	x.Write([]byte(str))
	y := x.Sum([]byte(""))

	fmt.Printf("%x\n", y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}
