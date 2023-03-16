package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(iters int) <-chan int {
	c := make(chan int)
	go func() {
		for i :=0; i < iters; i++ {
			c <- i
			time.Sleep(1 * time.Millisecond * 5)
		}
		close(c)
	}()
	return c
}

func consumer(cin <- chan int, sleepTime int) {
	for i := range cin {
		fmt.Println(i)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	}
}

func fanOut(ch <- chan int, size, lag int) []chan int {
	cs := make([]chan int ,size)
	for i, _ := range cs {
		cs[i] = make(chan int, lag)
	}
	
	go func() {
		var idx int64
		for i := range ch {
			chIdx := idx % int64(size)
			cs[chIdx] <- i
			idx++
		}
		for _, c := range cs {
			close(c)
		}
	}()
	return cs
}

func fanOutUnbuffered(ch <- chan int, size int) []chan int {
	cs := make([]chan int, size)
	for i, _ := range cs {
		cs[i] = make(chan int)
	}
	go func() {
		var idx int64
		for i := range ch {
			chIdx := idx % int64(size)
			cs[chIdx] <- i
			idx++
		}
		for _, c := range cs {
			close(c)
		}
	}()
	return cs
}




func main() {
	var consumerNum int = 10
	c := producer(1000)
	chans := fanOut(c, consumerNum, 10)
	var wg sync.WaitGroup
	for i := 0; i < consumerNum; i ++ {
		wg.Add(1)
		chn := chans[i]
		go func() {
			defer wg.Done()
			consumer(chn, 10)
		}()
	}
	wg.Wait()
}
