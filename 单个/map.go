package main

import "fmt"

func main() {
	//字典是无序的数据结构
	m := map[int]string{1: "mike", 2: "yoyo", 3: "go"}

	//第一个返回值为key ,第二个返回值为value,遍历结果是无序的
	for key, value := range m {
		fmt.Println("%d ========>%s", key, value)
	}

	//如何判断key值是否存在
	//第一个返回值为key所对应的value，第二个返回值为key是否存在的条件，存在ok为true
	value, ok := m[0]
	if ok == true {
		fmt.Println("m[1] = ", value)
	} else {
		fmt.Println("key不存在")
	}

	//删除map  delete(m,1)  第二个参数是key的值
	//函数调用的时候使用的是同一个 map
	//%+v  显示更详细
}
