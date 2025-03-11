package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err, time.Now(), "超时")
		return
	}

	// go outTime()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err, time.Now(), "连接失败")
			return
		}
		go process(conn)
	}
}

// 处理链接
func process(conn net.Conn) {
	addr := conn.RemoteAddr().String()
	fmt.Println("连接用户是", addr)
	var data [128]byte
	defer conn.Close()
	for {
		n, err := conn.Read(data[:])
		if err != nil {
			fmt.Println(err, time.Now(), "读取信息失败")
			return
		}
		fmt.Println(addr, string(data[:n]))
		_, err = conn.Write([]byte("已收到信息"))
		if err != nil {
			fmt.Println("信息回传失败")
		}
	}
}

// 心跳检测，检测是否在线
func checkOnline() {

}

// 获取信息保存进文件
func saveInfo() {}

// 连接超时时间
func outTime() {
	base, _ := time.Parse("15:04:05", "00:00:00") //基准时间
	var outTime string = "00:01:00"               //设定的超时时间
	timer, _ := time.Parse("15:04:05", outTime)
	t := timer.Sub(base)
	for {
		if t < 0 {
			return
		}
		time.Sleep(time.Second)
		t = t - time.Second
		t2 := base.Add(t)
		a := t2.Format("15:04:05")
		fmt.Println(a)
	}

}
