package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string
	Name string
	Addr string
}

var OnlineMap map[string]Client

var message = make(chan string)

func main() {

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err, time.Now(), "超时")
		return
	}

	go manager()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err, time.Now(), "连接失败")
			return
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr().String()

	cli := Client{
		C:    make(chan string),
		Name: addr,
		Addr: addr,
	}

	OnlineMap[addr] = cli
	go sendMessageToAddr(conn, cli)
	message <- makemessage(cli, "login")

	//处理是哪一位用户
	cli.C <- makemessage(cli, "i am here")

	isQuit := make(chan bool) //是否为主动退出
	isData := make(chan bool) //是否有数据
	//处理客户端发送的消息
	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Println(err, time.Now())
				return
			}
			//转发此内容
			msg := string(buf[:n])
			fmt.Println(len(msg), msg, "------------")
			//查询当前用户列表
			if len(msg) == 4 && msg == "who\n" {
				conn.Write([]byte("User list\n"))

				for _, tmp := range OnlineMap {
					msg := tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else if len(msg) >= 8 && msg[0:6] == "rename" {
				newName := strings.Split(msg, "|")[1]
				cli.Name = newName
				OnlineMap[addr] = cli
				conn.Write([]byte("修改成功"))
			} else {
				//转发
				message <- makemessage(cli, msg)
			}
			isData <- true
		}
	}()
	for {
		select {
		case <-isQuit:
			delete(OnlineMap, cli.Addr)
			message <- makemessage(cli, "退出")
			return
		case <-isData:
		case <-time.After(3 * time.Second):
			delete(OnlineMap, cli.Addr)
			message <- makemessage(cli, "超时退出 ")
			return
		}
	}
}
func makemessage(cli Client, str string) string {
	msg := "[" + cli.Addr + "]" + "[" + cli.Name + "]" + str + "\n"
	return msg
}
func sendMessageToAddr(conn net.Conn, cli Client) {
	for msg := range cli.C {
		conn.Write([]byte(msg))
	}
}

func manager() {
	OnlineMap = make(map[string]Client)
	for {
		msg := <-message
		for _, cli := range OnlineMap {
			cli.C <- msg
		}
	}
}

// func delete(oMap map[string]Client, cAddr string) {}
