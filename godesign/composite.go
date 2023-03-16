package main

import "fmt"

// Component 接口表示树形结构中的节点
type Component interface {
	Operation() string
}

// Leaf 结构体表示叶子节点
type Leaf struct{}

func (l *Leaf) Operation() string {
	return "Leaf"
}

// Composite 结构体表示组合节点
type Composite struct {
	children []Component
}

func (c *Composite) Operation() string {
	result := ""
	for _, child := range c.children {
		result += child.Operation() + "+"
	}
	return result + "Composite"
}

func (c *Composite) Add(component Component) {
	c.children = append(c.children, component)
}

func (c *Composite) Remove(component Component) {
	for i, child := range c.children {
		if child == component {
			c.children = append(c.children[:i], c.children[i+1:]...)
			break
		}
	}
}

// 示例
func main() {
	leaf := &Leaf{}
	
	composite1 := &Composite{}
	composite1.Add(leaf)
	composite1.Add(leaf)
	
	composite2 := &Composite{}
	composite2.Add(composite1)
	composite2.Add(leaf)
	
	composite3 := &Composite{}
	composite3.Add(composite2)
	composite3.Add(leaf)
	composite3.Add(leaf)
	
	fmt.Println(composite3.Operation())
}

