package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var CLOSEA = false
var DATA = make(map[int]bool)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
// 只写channel
func first(min, max int, out chan<- int) {
	for {
		if CLOSEA {
			fmt.Println("close send channel num")
			close(out)
			return
		}
		num := random(min, max)
		fmt.Printf("generator num %d\n", num)
		out<- num
	}
}

// 只读channel
func second(out chan<- int, in <-chan int) {
	for x := range in {
		fmt.Printf("get number %d\n", x)
		_, ok := DATA[x]
		if ok {
			fmt.Printf("duplidate num %d not send\n", x)
			CLOSEA = true
		} else {
			DATA[x] = true
			fmt.Printf("send num %d\n", x)
			out <- x
		}
	}
	fmt.Println("close send channel num")
	close(out)
}

// 只读channel
func third(in <-chan int) {
	var sum int
	sum = 0
	for x2 := range in {
		fmt.Printf("add num %d to sum\n", x2)
		sum = sum + x2
	}
	fmt.Printf("the sum of the random numbers is %d\n", sum)
}


func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need two int parameters")
		os.Exit(1)
	}
	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])
	
	if n1 > n2 {
		fmt.Printf("%d should be smaller then %d\n", n1, n2)
		return
	}
	
	rand.Seed(time.Now().UnixNano())
	A := make(chan int)
	B := make(chan int)
	go first(n1, n2, A)
	go second(B, A)
	third(B)
}