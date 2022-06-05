package main

func main() {
	defer func() {
		println("func1")
	}() //()括号的作用是调用匿名函数
	defer println("func2")
	println("main")
}
