package main

import (
	"testing"
)

func TestDecoratorFunc(t *testing.T) {
	a := new(DecorateA)
	b := new(DecorateB)
	c := new(DecorateC)
	b.Base.DecorateFun(a)
	c.Base.DecorateFun(b)
	c.Do()
}
