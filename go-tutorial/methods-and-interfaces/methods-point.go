package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}


func (v *Vertex) Scale (f float64) {
	v.X *= f
	v.Y *= f
}


func main () {
	vt := Vertex{3, 4}
	vt.Scale(2)
	fmt.Println("x:", vt.X)
	fmt.Println("y:", vt.Y)

}
