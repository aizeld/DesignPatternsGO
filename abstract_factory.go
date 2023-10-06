package main

import "fmt"

type AbstractFactory interface {
	CreateProductA() ProductA
	CreateProductB() ProductB
}

type ConcreteFactory1 struct{}

func (f *ConcreteFactory1) CreateProductA() ProductA {
	return &ConcreteProductA1{}
}

func (f *ConcreteFactory1) CreateProductB() ProductB {
	return &ConcreteProductB1{}
}

type ConcreteFactory2 struct{}

func (f *ConcreteFactory2) CreateProductA() ProductA {
	return &ConcreteProductA2{}
}

func (f *ConcreteFactory2) CreateProductB() ProductB {
	return &ConcreteProductB2{}
}

type ProductA interface {
	UseA() string
}

type ConcreteProductA1 struct{}

func (p *ConcreteProductA1) UseA() string {
	return "ProductA1"
}

type ConcreteProductA2 struct{}

func (p *ConcreteProductA2) UseA() string {
	return "ProductA2"
}

type ProductB interface {
	UseB() string
}

type ConcreteProductB1 struct{}

func (p *ConcreteProductB1) UseB() string {
	return "ProductB1"
}

type ConcreteProductB2 struct{}

func (p *ConcreteProductB2) UseB() string {
	return "ProductB2"
}

func main() {
	factory1 := &ConcreteFactory1{}

	productA1 := factory1.CreateProductA()
	productB1 := factory1.CreateProductB()

	fmt.Println(productA1.UseA())
	fmt.Println(productB1.UseB())

	factory2 := &ConcreteFactory2{}

	productA2 := factory2.CreateProductA()
	productB2 := factory2.CreateProductB()

	fmt.Println(productA2.UseA())
	fmt.Println(productB2.UseB())
}