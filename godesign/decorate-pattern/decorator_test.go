package main

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
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