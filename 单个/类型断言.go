package main

import "fmt"

type Student struct {
	name string
	id   int
}

//类型断言 if语句
func main() {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "hello go"
	i[2] = Student{"mike", 666}

	//类型查询，类型断言
	//for断言
	//第一个返回下标，第二个返回下标对应的值，data分别是i[0], i[1], i[2]
	for index, data := range i {
		//第一个返回的是值，第二个返回结果的真假
		if value, ok := data.(int); ok == true {
			fmt.Printf("x[%d] 类型为int, 内容为%d\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("x[%d] 类型为string, 内容为%s\n", index, value)
		} else if value, ok := data.(Student); ok == true {
			fmt.Printf("x[%d] 类型为Student, 内容为 name = %s, id = %d", index,
				value.name, value.id)
		}
	}

	//case 断言
	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Printf("x[%d] 类型为int, 内容为%d\n", index, value)
		case string:
			fmt.Printf("x[%d] 类型为string, 内容为%s\n", index, value)
		case Student:
			fmt.Printf("x[%d] 类型为Student, 内容为 name = %s, id = %d", index,
				value.name, value.id)
		}
	}

}
