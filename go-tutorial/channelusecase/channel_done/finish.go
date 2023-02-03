package main

import (
	"fmt"
)

type channel struct {
	in chan int
	done chan bool
}

func workGo(id int, c chan int, done chan bool) {
	for a := range c {
		fmt.Printf("work %d received %d\n", id, a)
		done <- true
	}
}

func work(id int) channel {
	c := channel{
		in:   make(chan int),
		done: make(chan bool),
	}
	go workGo(id, c.in, c.done)
	return c
}

func chanMake() {
	var channels [9]channel
	for i := 0; i < 9; i++ {
		channels[i] = work(i)
	}
	for i :=0; i<9;i++ {
		channels[i].in <- i
		<-channels[i].done
	}
}

func main() {
	chanMake()
	// time.Sleep(time.Millisecond * 10)
}