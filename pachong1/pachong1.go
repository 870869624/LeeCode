package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {

	var start, end int
	fmt.Println("请输入开始页(>=1)")
	fmt.Scanln(&start)
	fmt.Println("请输入终止页(>=1)")
	fmt.Scanln(&end)

	doWork(start, end)
	// resp, err := http.Get("http://www.baidu.com")
	// if err != nil {
	// 	fmt.Println(time.Now(), err)
	// 	return
	// }
	// fmt.Println(resp.Body, resp.Status)

	// buf := make([]byte, 4*1024)
	// var tmp string
	// for {
	// 	n, err := resp.Body.Read(buf)
	// 	if n == 0 {
	// 		fmt.Println(time.Now(), err)
	// 		return
	// 	}
	// 	tmp += string(buf[:n])
	// 	fmt.Println(tmp)
	// }
}

func doWork(start, end int) {
	page := make(chan int)
	for i := start; i <= end; i++ {
		go spiderPage(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("%d已完成", <-page)
	}
}

func HttpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	var result string
	buf := make([]byte, 5*1024)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("body", err)
			break
		}
		result += string(buf[:n])
	}
	return result, nil
}

func spiderPage(i int, page chan int) {
	str := strconv.Itoa((i - 1) * 50)
	url := "http://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + str
	fmt.Println(url)

	result, err := HttpGet(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileName := strconv.Itoa(i) + ".html"
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.Write([]byte(result))
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Close()
	page <- i
}
