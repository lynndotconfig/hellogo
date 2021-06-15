package go_tutorial

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("Hello")
}
