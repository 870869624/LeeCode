package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	var path string
	fmt.Println("输入要传输的文件名")
	fmt.Scanln(&path)
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	conn, err := net.Dial("tcp", ":8081")
	if err != nil {
		fmt.Println(err, time.Now(), "超时")
		return
	}
	defer conn.Close()

	for {
		bfo := bufio.NewReader(os.Stdin)
		input, _ := bfo.ReadString('\n')
		str := strings.Trim(input, "\r\n")
		if strings.ToUpper(str) == "Q" {
			_, err = conn.Write([]byte("用户已推出"))
			if err != nil {
				fmt.Println(err, time.Now(), "传输出现问题")
				return
			}

			fmt.Println("已推出")
			return
		}

		_, err = conn.Write([]byte(str))
		if err != nil {
			fmt.Println(err, time.Now(), "传输出现问题")
			return
		}
	}
}
