package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(time.Now(), err)
		return
	}

	for {
		conn, _ := listen.Accept()
		addr := conn.RemoteAddr().String()
		fmt.Println("已连接", addr)
		go process(conn)
		go sent(conn)
		go creatFile(conn)
	}
}

func process(conn net.Conn) {
}

func sent(conn net.Conn) {
	bfo := bufio.NewReader(os.Stdin)
	str, _ := bfo.ReadString('\n')
	defer conn.Close()
	conn.Write([]byte(str))

}

func creatFile(conn net.Conn) {
	defer conn.Close()
	var data [512]byte
	for {
		n, _ := conn.Read(data[:]) //读取文件路径
		conn.Write([]byte("OK"))
		fileName := string(data[:n])

		f, err := os.Create(fileName)
		if err != nil {
			fmt.Println(time.Now(), err)
			return
		}
		// defer f.Close()
		n, err = conn.Read(data[:]) //读取文件内容
		if err != nil {
			if err == io.EOF {
				fmt.Println("读取完毕")
			} else {
				fmt.Println(time.Now(), err)
				return
			}
		}
		if n == 0 {
			fmt.Println("读取完毕")
		}

		fmt.Println(string(data[:n]), fileName)

		_, err = f.Write(data[:n])
		if err != nil {
			fmt.Println(time.Now(), err)
			return
		}
		fmt.Println("q11")
	}
}
