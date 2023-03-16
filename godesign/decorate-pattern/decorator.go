package main

import (
	"fmt"
)

// 饮料接口
type Beverage interface {
	GetDescription() string
	Cost() float64
}

// 咖啡
type Coffee struct {}

func (c *Coffee) GetDescription() string {
	return "Coffee"
}

func (c *Coffee) Cost() float64 {
	return 2.0
}

// 茶
type Tea struct {}

func (t *Tea) GetDescription() string {
	return "Tea"
}

func (t *Tea) Cost() float64 {
	return 1.5
}

// 装饰器接口
type CondimentDecorator interface {
	GetDescription() string
	Cost() float64
}

// 牛奶装饰器
type Milk struct {
	beverage Beverage
}


func (m *Milk) GetDescription() string {
	return m.beverage.GetDescription() + ", Milk"
}

func (m *Milk) Cost() float64 {
	return m.beverage.Cost() + 0.5
}

// 糖装饰器
type Sugar struct {
	beverage Beverage
}

func (s *Sugar) GetDescription() string {
	return s.beverage.GetDescription() + ", Sugar"
}

func (s *Sugar) Cost() float64 {
	return s.beverage.Cost() + 0.3
}

func main() {
	var coffee Beverage
	coffee = &Coffee{}
	coffee = &Milk{coffee}
	coffee = &Sugar{coffee}
	
	fmt.Println(coffee.GetDescription()) // Output: Coffee, Milk, Sugar
	fmt.Println(coffee.Cost()) // Output: 3.0
	
	var tea Beverage
	tea = &Tea{}
	tea = &Milk{tea}
	tea = &Sugar{tea}
	
	fmt.Println(tea.GetDescription()) // Output: Coffee, Milk, Sugar
	fmt.Println(tea.Cost()) // Output: 3.0
}