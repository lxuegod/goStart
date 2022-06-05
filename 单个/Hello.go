package main

//基于图灵API 一个简单的聊天机器人
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

//请求结构体
type requestBody struct {
	Key    string `json:"key"`
	Info   string `json:"info"`
	UserId string `json:"userid"`
}

//结果体结构体
type responseBody struct {
	Code int      `json:"code"`
	Text string   `json:"string"`
	List []string `json:"list"`
	Url  string   `json:"url"`
}

//请求机器人
func process(inputChan <-chan string, userid string) {
	for {
		//从管道中接收输入
		input := <-inputChan
		if input == "EOF" {
			break
		}
		//构建结构体请求
		reqData := &requestBody{
			Key:    "792bcf45156d488c92e9d11da494b085",
			Info:   input,
			UserId: userid,
		}
		//转义为json
		byteData, _ := json.Marshal(&reqData)
		// if err != nil {
		// 	fmt.Println("json.Marshal error:", err)
		// 	return
		// }
		//请求聊天机器人接口
		req, err1 := http.NewRequest("POST",
			"http://www.tuling123.com/openapi/api",
			bytes.NewReader(byteData))
		if err1 != nil {
			fmt.Println("http.NewRequest error", err1)
			return
		}
		req.Header.Set("Content-Type", "application/json;charset=UTF-8")
		client := http.Client{}
		resp, err2 := client.Do(req)
		if err2 != nil {
			fmt.Println("http.NewRequest error", err2)
			return
		} else {
			//结果从json中解析并输出到命令行
			body, _ := ioutil.ReadAll(resp.Body)
			var respData responseBody
			json.Unmarshal(body, &respData)
			fmt.Println("AI: " + respData.Text)
		}
		if resp != nil {
			resp.Body.Close()
		}
	}
}

func main() {
	var input string
	fmt.Println("Enter 'EOF' to shut down :")
	//创建通道
	channel := make(chan string)
	//main函数结束时关闭通道
	defer close(channel)
	//启动goroutine运行机器人回答 线程
	go process(channel, string(rand.Int63()))

	for {
		//从命令行中读取输入
		fmt.Scanf("%s", &input)
		//将输入放到管道中
		channel <- input
		//结束程序
		if input == "EOF" {
			fmt.Println("Bye!")
			break
		} /* else if input == "你好" {
			fmt.Println("你陪我玩我就好啦")
		} else if input == "今天天气真好" {
			fmt.Println("是很好。没太阳")
		}*/
	}
}
