package main

import (
	"fmt"
	"net"
	"os"
)

//发送文件
func SendFile(path string, conn net.Conn) {
	//只读方式打开
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open err = ", err)
		return
	}

	defer f.Close()

	buf := make([]byte, 1024*4) //4k缓冲
	//读文件内容，读多少发多少，一点不差
	for {
		n, err := f.Read(buf)
		if err != nil {
			fmt.Println("os.Open err = ", err)
			return
		}

		//发送内容
		conn.Write(buf[:n])

	}
}

func main() {
	//if "connect" == os.Args[1]

	//提示输入文件
	fmt.Println("请输入需要传输的文件：")
	var path string
	fmt.Scan(&path)

	//获取文件名
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//主动连接服务器
	conn, err1 := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	defer conn.Close()

	//给接收方先发送文件名
	_, err = conn.Write([]byte(info.Name()))
	if err != nil {
		fmt.Println("conn.Write err = ", err)
		return
	}

	//接收对方的回复，如果对方回复"ok",说明对方已经准备好(对方可以接收文件)，可以发送文件
	var n int
	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err = ", err)
		return
	}

	if "ok" == string(buf[:n]) {
		fmt.Println("收到recive的回复ok")
		//发送文件内容
		SendFile(path, conn)
	}

}
