package main

import (
	"fmt"
	"strconv"
)

func fibonacci() func(int) int {
	var m = make(map[string]int)
	return func(x int) int {
		t := strconv.Itoa(x)
		k, ok := m[t]
		// todo
		}
}

func main() {
	f := bibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i)
	}
}	
