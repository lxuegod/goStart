package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func RecvFile(fileName string, conn net.Conn) {
	//新建文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err = ", err)
		return
	}

	buf := make([]byte, 4*1024)
	//接收多少，读多少，一点不差
	for {
		n, err := conn.Read(buf) //读取对方发送的文件名
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn.Read err = ", err)
				return
			}
		}
		if n == 0 {
			fmt.Println("n == 0 文件接收完毕")
			break
		}

		f.Write(buf[:n])
	}

}

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer listener.Close()

	//阻塞等待用户连接
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	buf := make([]byte, 1024)
	var n int
	n, err = conn.Read(buf) //读取对方发送的文件名
	if err != nil {
		fmt.Println("conn.Read err = ", err)
		return
	}
	fileName := string(buf[:n])

	//回复"ok"
	conn.Write([]byte("ok"))

	//接收文件内容
	RecvFile(fileName, conn)
}
