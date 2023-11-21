package main

import "fmt"

//strategy change

//strategy

type iattack interface {
	attack()
}

type slowAttack struct{}
type rapidAttack struct{}

func (r *rapidAttack) attack() {
	fmt.Println("Fires rapidly")
}
func (s *slowAttack) attack() {
	fmt.Println("fires slowly")
}

func (g *Gun) setStrategy(attackstrategy iattack) {
	g.attackStrategy = attackstrategy
}

func (g *Gun) getAttack() {
	g.attackStrategy.attack()
}

///end of strategy

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
	getAttack()
}

type Gun struct {
	name           string
	power          int
	attackStrategy iattack
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

type Ak47 struct {
	Gun
}

func newAk47(attackStrategy iattack) IGun {
	return &Ak47{
		Gun: Gun{
			name:           "AK47 gun",
			power:          4,
			attackStrategy: attackStrategy,
		},
	}
}

type Musket struct {
	Gun
}

func newMusket(attackStrategy iattack) IGun {
	return &Musket{
		Gun: Gun{
			name:           "Musket gun",
			power:          1,
			attackStrategy: attackStrategy,
		},
	}
}

type PowerIncreasingDecorator struct {
	IGun
	powerIncrease int
}

func (pid *PowerIncreasingDecorator) getPower() int {
	return pid.IGun.getPower() + pid.powerIncrease
}

func newPowerIncreasingDecorator(gun IGun, powerIncrease int) IGun {
	return &PowerIncreasingDecorator{
		IGun:          gun,
		powerIncrease: powerIncrease,
	}
}

func (pid *PowerIncreasingDecorator) setPower(power int) {
	pid.IGun.setPower(power)
}
func (pid *PowerIncreasingDecorator) setName(name string) {
	pid.IGun.setName(name)
}
func (pid *PowerIncreasingDecorator) getName() string {
	return pid.IGun.getName()
}

func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(&rapidAttack{}), nil
	}
	if gunType == "musket" {
		return newMusket(&slowAttack{}), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	ak47WithIncreasedPower := newPowerIncreasingDecorator(ak47, 5)

	printDetails(ak47)
	printDetails(musket)
	printDetails(ak47WithIncreasedPower)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()

	g.getAttack()
}
