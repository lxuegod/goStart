//服务器socket
package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

var port = ":8888"

type Client struct {
	C    chan string //用户发送数据的管道
	Name string      //用户名
	Addr string      //网络地址
}

//保存在线用户    cliAddr======>Client
var onlieMap map[string]Client

//广播用的信息
var message = make(chan string)

//新开一个协程，转发消息，只要有消息来了，遍历map，给map的所有成员都发送此消息
func Manger() {
	//给map分配空间
	onlieMap = make(map[string]Client)
	for {
		msg := <-message //没有消息前，这里会阻塞

		//遍历map，给map的所有成员都发送信息
		for _, cli := range onlieMap {
			cli.C <- msg
		}
	}
}

//给当前客户端发信息
func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C { //给当前客户端发信息
		conn.Write([]byte(msg + "\n"))
	}
}

//编辑信息
func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]," + cli.Name + ":\n" + msg
	return

}

//处理用户连接
func HandleConn(conn net.Conn) { //处理用户连接
	defer conn.Close()

	//获取客户端的网络地址
	cliAddr := conn.RemoteAddr().String()

	//创建一个结构体，默认用户名和网络地址一样
	cli := Client{make(chan string), cliAddr, cliAddr}
	//把结构体添加到map中
	onlieMap[cliAddr] = cli

	//新开一个协程，专门给当前客户端发信息
	go WriteMsgToClient(cli, conn)

	//广播某个在线
	message <- MakeMsg(cli, "login")
	//提示我是谁
	cli.C <- MakeMsg(cli, "I am come here")

	isQuit := make(chan bool)  //对方是否主动退出
	hasData := make(chan bool) //对方是否有数据发送

	//新建一个协程，接收用户发送的信息
	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if n == 0 { //对方断开或者出问题
				isQuit <- true
				fmt.Println("conn.Read err = ", err)
				return
			}

			//通过Windows命令行测试 多了两个字符\r\n  \r为转义字符
			msg := string(buf[:n-2]) //通过windows nc 测试多一个换行
			fmt.Println(len(msg))
			if len(msg) == 3 && msg == "who" {
				//遍历map给所有的用户发送所有成员
				conn.Write([]byte("\nUser list:\n"))
				for _, tmp := range onlieMap {
					msg := tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				onlieMap[cliAddr] = cli
				conn.Write([]byte("remane ok\n"))
			} else { //转发此内容
				message <- MakeMsg(cli, msg)
			}

			hasData <- true //代表有数据

		}
	}()

	for {
		//通过select监测channel的流动
		select {
		case <-isQuit:
			delete(onlieMap, cliAddr)            //当前用户从map中移除
			message <- MakeMsg(cli, "login out") //广播谁下线了
			return

		case <-hasData:

		case <-time.After(5 * time.Minute): //5分钟超时退出
			delete(onlieMap, cliAddr)                     //当前用户从map中移除
			message <- MakeMsg(cli, "time out leave out") //广播谁下线了
			return
		}

	}
}

func main() {
	//监听
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	defer listener.Close()

	//新开一个协程，转发消息，只要有消息来了，遍历map，给map的所有成员都发送此消息
	go Manger()

	//主协程，循环阻塞等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err = ", err)
			continue
		}

		defer conn.Close()

		//处理用户连接
		go HandleConn(conn)

	}
}
