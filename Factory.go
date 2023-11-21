// BAD BAD BAAAAAAAAAD example of factory usage, i should be using newKhabib newJonJones instead of newWrestler or newStriker

package main

import "fmt"

//strategy

//Command

type ChangeHeight struct {
	fighter iFighter
	amount  int
	backup  int
}

func newChangeHeight(fighter iFighter, amount int) *ChangeHeight {
	return &ChangeHeight{
		fighter: fighter,
		amount:  amount,
		backup:  fighter.getHeight(),
	}
}

func (c *ChangeHeight) Execute() {
	c.fighter.setHeight(c.amount)
}
func (c *ChangeHeight) Undo() {
	c.fighter.setHeight(c.backup)
}

//decorator

type SpecialFighter struct {
	fighter iFighter
	persona string
}

func (tf *SpecialFighter) getHeight() int {
	return tf.fighter.getHeight()
}

func (tf *SpecialFighter) setHeight(height int) {
	tf.fighter.setHeight(height)
}

func (tf *SpecialFighter) getName() string {
	return tf.fighter.getName() + tf.persona
}
func (tf *SpecialFighter) setName(name string) {
	tf.fighter.setName(name)
}
func (tf *SpecialFighter) getReach() int {
	return tf.fighter.getReach()
}
func (tf *SpecialFighter) setReach(reach int) {
	tf.fighter.setReach(reach)
}

func ChangePersona(fighter iFighter, persona string) iFighter {
	return &SpecialFighter{
		fighter: fighter,
		persona: persona,
	}
} //END OF DECORATOR

// it doesnt change the fighters height, but it decorates, so logically it's right implementation
type fighter struct {
	name   string
	height int
	reach  int
}

type iFighter interface {
	setName(name string)
	getName() string
	setHeight(height int)
	getHeight() int
	setReach(reach int)
	getReach() int
}

func (f *fighter) setName(name string) {
	f.name = name
}
func (f *fighter) getName() string {
	return f.name
}

func (f *fighter) setHeight(height int) {
	f.height = height
}
func (f *fighter) getHeight() int {
	return f.height
}

func (f *fighter) setReach(reach int) {
	f.reach = reach
}
func (f *fighter) getReach() int {
	return f.reach
}

type wrestler struct {
	fighter
}
type striker struct {
	fighter
}

func newWrestler(name string) iFighter {
	return &wrestler{
		fighter: fighter{
			name:   name,
			reach:  65,
			height: 175,
		},
	}
}

func newStriker(name string) iFighter {
	return &striker{
		fighter: fighter{
			name:   name,
			reach:  85,
			height: 188,
		},
	}
}

func GetFighters(style string, name string) (iFighter, error) {
	if style == "wrestling" {
		return newWrestler(name), nil
	}
	if style == "striker" {
		return newStriker(name), nil
	}
	return nil, fmt.Errorf("SHIT WHO THE FUK IS THAT")
}

func Printstats(f iFighter) {
	fmt.Printf("name: %s", f.getName())
	fmt.Println()
	fmt.Printf("reach: %d", f.getReach())
	fmt.Println()
	fmt.Printf("height: %d", f.getHeight())
	fmt.Println()
}

func main() {
	Khabib, _ := GetFighters("wrestling", "khabibi")

	KhabibWithHat := ChangePersona(Khabib, " with PAPAHA")

	shavkat, _ := GetFighters("striker", "Shavkat Rakhmonov")

	shavkat = ChangePersona(shavkat, " With Tymak")

	ChangeHeight := newChangeHeight(shavkat, 186)
	ChangeHeight.Execute()
	ChangeHeight.Undo()
	Sergei, _ := GetFighters("striker", "Sergei Pavlovich")

	Printstats(KhabibWithHat)
	Printstats(Sergei)
	Printstats(shavkat)
}
