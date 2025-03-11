package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

func main() {
	var start, end int
	fmt.Println("请输入开始页(>=1)")
	fmt.Scanln(&start)
	fmt.Println("请输入终止页(>=1)")
	fmt.Scanln(&end)

	doWork(start, end)
}

func doWork(strat, end int) {
	fmt.Printf("%d页到%d页", strat, end)
	for i := strat; i <= end; i++ {
		spiderPage(i)
	}
}

func spiderPage(i int) {

	url := "http://www.xiaokuwang.com/2-" + strconv.Itoa(i) + ".html"
	fmt.Println("正在爬取第%d", i, url)

	result, err := httpGet(url)
	if err != nil {
		fmt.Println("获取失败", url)
		return
	}

	fmt.Println(result)

	re1 := regexp.MustCompile(`<a href="(?s:(.*?))"`)
	if re1 == nil {
		fmt.Println("错误")
		return
	}
	str := re1.FindAllStringSubmatch(result, -1)
	fmt.Println("1111111111111", str, "11111111111111111111")
}

func httpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("获取失败", url)
		return "", err
	}
	defer resp.Body.Close()

	buf := make([]byte, 5*1024)
	n, err := resp.Body.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf[:n]), nil
}
