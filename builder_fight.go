package main

import "fmt"

type fighter struct {
	striking        int
	grappling       int
	mentalToughness int
}

type iFightBuilder interface {
	setStriking()
	setGrappling()
	setMindset()
	makeFighter() fighter
}

//Khabib

type khabib struct {
	striking        int
	grappling       int
	mentalToughness int
}

func newKhabib() *khabib {
	return &khabib{}
}

func (k *khabib) setStriking() {
	k.striking = 7
}
func (k *khabib) setGrappling() {
	k.grappling = 10
}
func (k *khabib) setMindset() {
	k.mentalToughness = 10
}

func (k *khabib) makeFighter() fighter {
	return fighter{
		striking:        k.striking,
		grappling:       k.grappling,
		mentalToughness: k.mentalToughness,
	}
}

//Coach

type coach struct {
	fighter iFightBuilder
}

func newCoach(f iFightBuilder) *coach {
	return &coach{
		fighter: f,
	}
}

func (c *coach) setCoach(f iFightBuilder) {
	c.fighter = f
}

func (c *coach) Coach() fighter {
	c.fighter.setGrappling()
	c.fighter.setStriking()
	c.fighter.setMindset()
	return c.fighter.makeFighter()
}

func getFighter(name string) (iFightBuilder, error) {
	if name == "khabiba" {
		return newKhabib(), nil
	}
	return nil, fmt.Errorf("Not found")
}

func main() {
	khabib, _ := getFighter("khabiba")
	Coach := newCoach(khabib)
	Javier := Coach.Coach()
	fmt.Println("Grappling level ", Javier.grappling)
	fmt.Println("Striking level ", Javier.striking)
	fmt.Println("Mental toughness ", Javier.mentalToughness)

}
