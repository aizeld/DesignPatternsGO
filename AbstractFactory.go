package main

import "fmt"

//Command pattern

type Command interface {
	Execute()
}

type ChangeShoes struct {
	shoe    IShoe
	newSize int
	backup  int
}

func newChangeShoes(shoe IShoe, newsize int) *ChangeShoes {
	return &ChangeShoes{
		shoe:    shoe,
		newSize: newsize,
		backup:  shoe.getSize(),
	}
}

func (n *ChangeShoes) Execute() {
	n.shoe.setSize(n.shoe.getSize() + n.newSize)
}

func (n *ChangeShoes) Undo() {
	n.shoe.setSize(n.backup)
}

/// END  OF COMMAND

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printJeansDetails(s IJeans) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func main() {
	adidasFactory, _ := getShit("adidas")
	adidasShoe := adidasFactory.makeShoes(15)

	Changeshoe := newChangeShoes(adidasShoe, 12)
	Changeshoe.Execute()
	adidasJeans := adidasFactory.makeJeans(48)

	printJeansDetails(adidasJeans)
	printShoeDetails(adidasShoe)
}

type adidas struct{}
type nike struct{}

type isportsfactory interface {
	makeJeans(size int) IJeans
	makeShoes(size int) IShoe
}

func getShit(logo string) (isportsfactory, error) {

	if logo == "adidas" {
		return &adidas{}, nil
	}
	if logo == "nike" {
		return &nike{}, nil
	}
	return nil, fmt.Errorf("USUALLY SHITE DIVE TO CLOSE THE DISTANCE")

}

// / only jeans
type Jeans struct {
	size int
	logo string
}
type IJeans interface {
	setLogo(logo string)
	getLogo() string
	setSize(size int)
	getSize() int
}

func (s *Jeans) setLogo(logo string) {
	s.logo = logo
}
func (s *Jeans) getLogo() string {
	return s.logo
}

func (s *Jeans) setSize(size int) {
	s.size = size
}
func (s *Jeans) getSize() int {
	return s.size
}

type adidasJeans struct {
	Jeans
}
type nikeJeans struct {
	Jeans
}

func (a *adidas) makeJeans(size int) IJeans {
	return &adidasJeans{
		Jeans: Jeans{
			logo: "adidas",
			size: size,
		},
	}
}

func (a *nike) makeJeans(size int) IJeans {
	return &nikeJeans{
		Jeans: Jeans{
			logo: "nike",
			size: size,
		},
	}
}

// /Only shoes
type Shoe struct {
	size int
	logo string
}

type IShoe interface {
	setLogo(logo string)
	getLogo() string
	setSize(size int)
	getSize() int
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}
func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}
func (s *Shoe) getSize() int {
	return s.size
}

type adidasShoes struct {
	Shoe
}
type nikeShoes struct {
	Shoe
}

func (a *adidas) makeShoes(size int) IShoe {
	return &adidasShoes{
		Shoe: Shoe{
			logo: "adidas",
			size: size,
		},
	}
}
func (a *nike) makeShoes(size int) IShoe {
	return &nikeShoes{
		Shoe: Shoe{
			logo: "nike",
			size: size,
		},
	}
}
