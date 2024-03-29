package go_tutorial

import "fmt"

const (
	// Create a huge number by shifting a  1 bit left 100 place.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 place, so we end up with 1 << 2, or 2
	Small = Big >> 99
)

func needInt(x int) int { return x * 10 + 1 }

func needFloat(x float64) float64 { return x * 0.1}

func main () {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
