package main

import (
	"fmt"
)

type Component interface {
	Operation() string
}

type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
	return "ConcreteComponent"
}


type Decorator struct {
	component Component
}

func (d *Decorator) Operation() string {
	return d.component.Operation()
}

type ConcreteDecoratorA struct {
	Decorator
}

func (c *ConcreteDecoratorA) Operation() string {
	return "ConcreteDecoratorA(" + c.component.Operation() + ")"
}

type ConcreteDecoratorB struct {
	Decorator
}

func (c *ConcreteDecoratorB) Operation() string {
	return "ConcreteDecoratorB(" + c.component.Operation() + ")"
}

func main() {
	component := &ConcreteComponent{}
	decoratorA := &ConcreteDecoratorA{component}
	decoratorB := &ConcreteDecoratorB{decoratorA}
	fmt.Println("ConcreteComponent: " + component.Operation())
	fmt.Println("DecoratorA: " + decoratorA.Operation())
	fmt.Println("DecoratorB: " + decoratorB.Operation())
}