package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// method
func (v  Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// function
func Absf(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	vtx := Vertex{3,4}
	fmt.Println("methods is: ", vtx.Abs())
	fmt.Println("functions is: ", Absf(vtx))
}
