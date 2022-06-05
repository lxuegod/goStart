package main

import (
	"fmt"
	"regexp"
)

func main() {

	buf := "abc azc a7c aac 888 a9c tac"
	//1)解释规则，它会解析正则表达式，如果成功返回解释器

	//一般根据需要查找文档
	//reg1 := regexp.MustCompile(`a.c`)
	//reg1 := regexp.MustCompile(`a[0-9]c`)  另一种表示方法
	reg1 := regexp.MustCompile(`a\dc`)

	if reg1 == nil { //解析失败 返回nil
		fmt.Println("regexp error")
		return
	}

	//2)根据规则提取关键信息
	result1 := reg1.FindAllStringSubmatch(buf, -1) //-1为全部 其他数字为几个
	fmt.Println("result1 = ", result1)

	//另一种灵活的方法
	buf1 := "43.14 567 agsdg 1.23 7. 8.9 1sdljgl 6.66 7.8 "
	//reg := regexp.MustCompile(`\d.\d`)  \d是一个数字
	//\d   +匹配前一个字符的一个或多个
	reg := regexp.MustCompile(`\d+\.\d+`)
	if reg == nil {
		fmt.Println("regexp error")
		return
	}
	//result := reg.FindAllString(buf1, -1)
	result := reg.FindAllStringSubmatch(buf1, -1)
	fmt.Println("result = ", result)

}
