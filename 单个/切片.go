package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	s1 := slice[:3] //	从下标为 0 的数字切到下标为 3 的数字（不包括3）
	fmt.Println(s1)
	s2 := slice[3:] //	从下标为3的数字切到结尾
	fmt.Println(s2)
	s3 := slice[1:4] //	从下标为1的数字切到下标为4的数字
	fmt.Println(s3)
	//用append函数对切片进行扩容
}
