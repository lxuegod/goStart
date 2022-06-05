//client
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

var ip = "192.168.1.110:8888"

func main() {
	//主动连接服务器
	//net.Dial第一个参数 "tcp/udp"  第二个参数 ip:端口
	conn, err := net.Dial("tcp", ip)
	addr := conn.RemoteAddr().String()
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	} else {
		fmt.Println(addr, "connect successfully")
	}

	//main函数调用完毕，关闭连接
	defer conn.Close()

	go func() {
		//从键盘输入内容，给服务器发送内容
		str := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(str) //从键盘读取内容
			if err != nil {
				fmt.Println("os.Stdin.Read err = ", err)
				return
			}
			//把输入的内容给服务器发送
			conn.Write(str[:n])

		}

	}()
	//接收服务器回复的数据
	buf := make([]byte, 1024) //切片缓存
	for {
		n, err := conn.Read(buf) //接收服务器的内容
		if err != nil {
			fmt.Println("conn.Read err = ", err)
			return
		}
		//打印接收到的内容[%v]%v

		now := time.Now()
		formatNow := now.Format("2006-01-02 15:04:05")
		fmt.Printf("%v  %v", formatNow, string(buf[:n]))
	}
}
