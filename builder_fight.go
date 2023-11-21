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

//Mcgregor
type Conor struct {
	striking        int
	grappling       int
	mentalToughness int
}

func newConor() *Conor {
	return &Conor{}
}

func (k *Conor) setStriking() {
	k.striking = 9
}
func (k *Conor) setGrappling() {
	k.grappling = 0
}
func (k *Conor) setMindset() {
	k.mentalToughness = 7
}

func (k *Conor) makeFighter() fighter {
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
	if name == "conor" {
		return newConor(), nil
	}
	return nil, fmt.Errorf("Not found")
}

///facade pattern
type FighterFacade struct {
	khabib iFightBuilder
	conor  iFightBuilder
}

func newFighterFacade() *FighterFacade {
	return &FighterFacade{
		khabib: newKhabib(),
		conor:  newConor()}
}
func (ff *FighterFacade) MakeKhabib() fighter {
	khabib, _ := getFighter("khabiba")
	coach := newCoach(khabib)
	fighter := coach.Coach()
	return fighter
}

func (ff *FighterFacade) MakeConor() fighter {
	conor, _ := getFighter("conor")
	coach := newCoach(conor)
	fighter := coach.Coach()
	return fighter
}

func (ff *FighterFacade) showFighter(fighter fighter) {
	fmt.Println("Grappling level ", fighter.grappling)
	fmt.Println("Striking level ", fighter.striking)
	fmt.Println("Mental toughness ", fighter.mentalToughness)
}

func main() {
	// khabib, _ := getFighter("khabiba")
	// Coach := newCoach(khabib)
	// conor, _ := getFighter("conor")
	// Coach.setCoach(conor) /// в чем прикол этого сеткоача?
	// Javier := Coach.Coach()

	// fmt.Println("Grappling level ", Javier.grappling)
	// fmt.Println("Striking level ", Javier.striking)
	// fmt.Println("Mental toughness ", Javier.mentalToughness)

	FighterFacade := newFighterFacade()
	khabib := FighterFacade.MakeKhabib()
	conor := FighterFacade.MakeConor()
	FighterFacade.showFighter(khabib)
	FighterFacade.showFighter(conor) // только что обьеденил фасад и билдер

}
