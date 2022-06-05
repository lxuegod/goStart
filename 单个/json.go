package main

import (
	"encoding/json"
	"fmt"
)

//成员变量首字母必须大写  通过结构体的方法实现
// type IT struct {
// 	Company  string
// 	Subjects []string
// 	IsOk     bool
// 	Price    float64
// }

type IT struct { //首字母变为小写的方法  二次编码
	Company  string   `json:"company"`
	Subjects []string `json:"-"`       //此字段不会输出到屏幕
	IsOk     bool     `json:",string"` //转化成字符串
	Price    float64
}

//通过map的方法实现
func Map() {
	m := make(map[string]interface{}, 4)
	m["company"] = "itcast"
	m["subjects"] = []string{"Go", "C++", "Python", "Test"}
	m["isok"] = true
	m["price"] = 666.666
	//格式化编码成json
	buf, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("2result = ", string(buf))
}

func Str() {
	//定义一个结构体变量 同时初始化
	s := IT{"itcast", []string{"Go", "C++", "Python", "Test"}, true, 666.66}
	//编码，根据内容生成json文本

	//普通的编码
	//buf, err := json.Marshal(s)
	//格式化编码
	buf, err := json.MarshalIndent(s, "", "	")

	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("1buf = ", string(buf))
}

func main() {
	Str()
	Map()
}
