package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int, 0)
	
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()] += 1
	}
	
	// ctl + D = EOF， 退出当前输入
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("line=%s, count=%d\n", line, n)
		}
	}
}