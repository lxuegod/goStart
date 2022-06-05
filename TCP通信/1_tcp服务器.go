package main

import (
	"fmt"
	"net"
	//"strings"
)

func HandleConn(conn net.Conn) {
	//函数调用完毕，自动关闭
	defer conn.Close()

	//获取客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, "connect sucessfully")

	buf := make([]byte, 2048) //缓存区2048

	for {
		//读取用户信息
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err = ", err)
			return
		}
		//打印读到的内容
		fmt.Printf("[%s]%s\n", addr, string(buf[:n]))
		//fmt.Println("len = ", len(string(buf[:n])))

		//if "exit" == string(buf[:n-1]) {nc 测试
		if "exit" == string(buf[:n-2]) { //自己写的客户端测试，发送时多了两个字符"\r\n"
			fmt.Println(addr, "exit")
			return
		}

		// //把内容转换成大写，再给用户发送
		// conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
		//读到什么，写什么
		conn.Write([]byte(string(buf[:n])))
	}
}

func main() {
	//监听
	//net.Listen使用方法  第一个参数"tcp/udp"  第二个参数 ip:端口
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	defer listener.Close()

	//接收多个用户
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err = ", err)
			return
		}

		//处理用户请求
		go HandleConn(conn)
	}
}
