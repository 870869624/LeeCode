package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	// 打印1的goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			time.Sleep(1 * time.Second)
			fmt.Print(1)
			ch1 <- struct{}{}
		}
		close(ch1)
	}()

	// 打印2的goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			<-ch1
			fmt.Print(2)
			ch2 <- struct{}{}
		}
		close(ch2)
	}()

	// 打印3的goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			<-ch2
			fmt.Println(3)
		}
	}()

	wg.Wait()
}
