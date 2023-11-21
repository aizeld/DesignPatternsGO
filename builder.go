package main

import "fmt"

type House struct {
	doortype   string
	floor      int
	windowtype string
}

type IBuilder interface {
	setDoor()
	setWindow()
	setFloor()
	getHouse() House
}

//rich house

type richBuilder struct {
	windowtype string
	doortype   string
	floor      int
}

func NewrichBuilder() *richBuilder {
	return &richBuilder{}
}

func (r *richBuilder) setDoor() {
	r.doortype = "Silver"
}
func (r *richBuilder) setFloor() {
	r.floor = 5
}
func (r *richBuilder) setWindow() {
	r.windowtype = "gold"
}

func (r *richBuilder) getHouse() House {
	return House{
		doortype:   r.doortype,
		floor:      r.floor,
		windowtype: r.windowtype,
	}
}

// shitbuilder

type shitbuilder struct {
	doortype   string
	floor      int
	windowtype string
}

func newshitbuilder() *shitbuilder {
	return &shitbuilder{}
}

func (s *shitbuilder) setDoor() {
	s.doortype = "Air"
}
func (s *shitbuilder) setFloor() {
	s.floor = 0
}
func (s *shitbuilder) setWindow() {
	s.windowtype = "sticks"
}
func (s *shitbuilder) getHouse() House {
	return House{
		doortype:   s.doortype,
		windowtype: s.windowtype,
		floor:      s.floor,
	}
}

//Nachalniko mana mana

type Nasalnik struct {
	builder IBuilder
}

func newNasalnik(b IBuilder) *Nasalnik {
	return &Nasalnik{
		builder: b,
	}
}

func (n *Nasalnik) setBuilder(b IBuilder) {
	n.builder = b
}

func (n *Nasalnik) BuildHouse() House {
	n.builder.setDoor()
	n.builder.setFloor()
	n.builder.setWindow()
	return n.builder.getHouse()
}

func getBuilder(level string) (IBuilder, error) {
	if level == "shite" {
		return newshitbuilder(), nil
	}
	if level == "rich" {
		return NewrichBuilder(), nil
	}
	return nil, fmt.Errorf("MOTHAFUKA")

}

type FighterFacade struct {
	khabib iFightBuilder
	conor  iFightBuilder
}

func main() {
	shitbuilder, _ := getBuilder("rich")
	director := newNasalnik(shitbuilder)
	richHouse := director.BuildHouse()

	fmt.Printf("Normal House Door Type: %s\n", richHouse.doortype)
	fmt.Printf("Normal House Window Type: %s\n", richHouse.windowtype)
	fmt.Printf("Normal House Num Floor: %d\n", richHouse.floor)

}
