package main

import "fmt"

//brick
type brick struct {
	width  int
	height int
}

//peg
type peg struct {
	radius int
	height int
}

func (p *peg) getRadius() int {
	return p.radius
}

//Interface

type Ipeg interface {
	getRadius() int
}

//adapter

type adapterBricks struct {
	brick *brick
}

func (a *adapterBricks) getRadius() int {
	return a.brick.width / 2
}

//hole

type hole struct {
	radius int
}

func (h *hole) fits(peg Ipeg) bool {
	return h.radius >= peg.getRadius()
}

func main() {
	hole := &hole{radius: 10}

	peg := &peg{radius: 11}

	//adapter

	brick := &brick{width: 18, height: 1}
	adapterBricks := &adapterBricks{brick: brick}
	hole.getRes(peg)
	hole.getRes(adapterBricks)
}
func (hole *hole) getRes(peg Ipeg) {
	if hole.fits(peg) {
		fmt.Println("it fits")
	} else {
		fmt.Println("TOO big")
	}
}
