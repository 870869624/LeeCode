package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func main() {

	wg.Add(3)
	go outPut(1)
	go outPut(2)
	go outPut(3)
	wg.Wait()
}

func outPut(num int) {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		mutex.Lock()
		time.Sleep(1 * time.Second)
		fmt.Printf("第%d个goroutine,第%d次执行\n", num, i)
		mutex.Unlock()
	}
}
