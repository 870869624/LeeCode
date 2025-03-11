package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type shulie struct {
	i int
	j int
	v int
}

func main() {
	var chessmap1 []shulie

	var str string = "go, go, hello,word"
	str = strings.Replace(str, "go", "aaa", 1)
	fmt.Println(str)

	var a [10][10]int
	a[2][2] = 10
	a[5][5] = 17
	//将结果展示出来
	for _, v := range a {
		for _, v1 := range v {
			fmt.Printf("  %d", v1)
		}
		fmt.Println()
	}
	//创建一个文件
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	//将结果放入稀疏数列,只存入行列以及值信息
	for i, v2 := range a {
		for j, v3 := range v2 {
			if v3 != 0 {
				shulie := shulie{
					i: i,
					j: j,
					v: v3,
				}
				chessmap1 = append(chessmap1, shulie)
				fmt.Println(shulie)
			}
		}
	}

	fmt.Println(chessmap1)
	//将稀疏数组中的值取出来
	var b [11][11]int
	for _, v := range chessmap1 {
		b[v.i][v.j] = v.v
	}

	n, err := writer.WriteString(fmt.Sprint(b))
	if err != nil {
		fmt.Println(err)
		return
	}
	writer.Flush()
	fmt.Println(n)
	file.Seek(0, 0)

	for _, v := range b {
		for _, v1 := range v {
			fmt.Printf("  %d", v1)
		}
		fmt.Println()
	}
	data := bufio.NewReader(file)

	v, _ := data.ReadString(0)
	str1 := strings.Split(v, "")
	fmt.Println(str1)

	fmt.Println(v)
}
