package main

import (
	"fmt"
)

type IDecorate interface {
	Do()
}

// 装饰器。实现接口，又定义了自己的事件DecorateFun，相当于抽象类
type Decorate struct {
	// 待装饰的抽象类
	decorate IDecorate
}

func (c *Decorate) DecorateFun(i IDecorate) {
	c.decorate = i
}

func (c *Decorate) Do() {
	if c.decorate != nil {
		c.decorate.Do()
	}
}

// 具体A装饰
type DecorateA struct {
	Base Decorate
}

// 重写方法，隐式实现接口
func (c *DecorateA) Do() {
	fmt.Printf("执行A装饰\n")
	c.Base.Do()
}

// 具体B装饰
type DecorateB struct {
	Base Decorate
}

// 重写方法，隐式实现接口
func (c *DecorateB) Do() {
	fmt.Printf("执行B装饰\n")
	c.Base.Do()
}

// 具体C装饰
type DecorateC struct {
	Base Decorate
}

// 重写方法，隐式实现接口
func (c *DecorateC) Do() {
	fmt.Printf("执行C装饰\n")
	c.Base.Do()
}

