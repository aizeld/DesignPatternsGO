package main

import "fmt"

//pizza

type Ipizza interface {
	GetPrice() int
}

type pizza struct {
}

func (p *pizza) GetPrice() int {
	return 20
}

type CheesePizza struct {
	pizza Ipizza
}

func (p *CheesePizza) GetPrice() int {
	return p.pizza.GetPrice() + 10
}

/// Printer

type IPrinter interface {
	print() string
}

type simplePrinter struct{}

func (s *simplePrinter) print() string {
	return "Shit WOW"
}

func DecoratorBold(p IPrinter) IPrinter {
	return PrinterFunc(func() string {
		return "<B>" + p.print() + "<B>"
	})
}

type PrinterFunc func() string

func (pf PrinterFunc) print() string {
	return pf()
}

func main() {
	pizza := &pizza{}
	WIthCheese := &CheesePizza{pizza}

	fmt.Println(WIthCheese.GetPrice())
}
