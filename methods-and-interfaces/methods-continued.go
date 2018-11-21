package main

import (
	"fmt"
)

type Myfloat float64


func Abs(f Myfloat) float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

func main () {
	myFloat := Myfloat(1.1)
	fmt.Println("result is: ", Abs(myFloat))
} 
