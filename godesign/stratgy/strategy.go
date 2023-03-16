package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

type PricingStrategy interface {
	CalculatePrice(product *Product) float64
}

type NormalPricing struct{}

func (s *NormalPricing) CalculatePrice(product *Product) float64 {
	return product.Price
}

type SalePricing struct{}

func (s *SalePricing) CalculatePrice(product *Product) float64 {
	return 0.9 * product.Price
}

type VIPPricing struct{}

func (s *VIPPricing) CalculatePrice(product *Product) float64 {
	return 0.8 * product.Price
}

func main() {
	// 创建一些商品
	products := []*Product{
		{"Product A", 100},
		{"Product B", 200},
		{"Product C", 300},
	}
	
	// 创建三种不同的价格策略
	normalPricing := &NormalPricing{}
	salePricing := &SalePricing{}
	vipPricing := &VIPPricing{}
	
	// 计算商品的价格，并打印结果
	for _, p := range products {
		var pricingStrategy PricingStrategy
		
		// 根据商品的名称来选择不同的价格策略
		switch p.Name {
		case "Product A":
			pricingStrategy = normalPricing
		case "Product B":
			pricingStrategy = salePricing
		case "Product C":
			pricingStrategy = vipPricing
		}
		
		price := pricingStrategy.CalculatePrice(p)
		fmt.Printf("%s: %.2f\n", p.Name, price)
	}
}

