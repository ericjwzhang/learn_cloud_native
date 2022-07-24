package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// 三个生产者， 两个消费者
var flag = false

func producer(threadID int, wg *sync.WaitGroup, ch chan string) {
	count := 0
	for !flag {
		time.Sleep(time.Second * 1)
		count++
		data := strconv.Itoa(threadID) + "---" + strconv.Itoa(count)
		fmt.Printf("procuder, %s\n", data)
		ch <- data
	}
	wg.Done()
}

func consumer(wg *sync.WaitGroup, ch chan string) {
	for data := range ch {
		time.Sleep(1 * time.Second)
		fmt.Printf("consumer, %s\n", data)
	}
	wg.Done()
}

func main() {
	chanStream := make(chan string, 10)
	wgPd := new(sync.WaitGroup)
	wgCs := new(sync.WaitGroup)

	for i := 0; i < 3; i++ {
		wgPd.Add(1)
		go producer(i, wgPd, chanStream)
	}

	for j := 0; j < 2; j++ {
		wgCs.Add(1)
		go consumer(wgCs, chanStream)
	}

	go func() {
		time.Sleep(5 * time.Second)
		flag = true
	}()

	wgPd.Wait()
	close(chanStream)
	wgCs.Wait()
}
