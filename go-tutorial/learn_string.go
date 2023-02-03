package go_tutorial

import (
	"fmt"
	"unicode/utf8"
)

func PrintLearnString() {
	s := "hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
}

