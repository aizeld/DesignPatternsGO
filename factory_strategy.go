package main

import (
	"fmt"
)

//attack strategy

type iattack interface {
	attack() string
}
type slowAttack struct{}

func (s *slowAttack) attack() string {
	return "Fires slow"
}

type rapidAttack struct{}

func (s *rapidAttack) attack() string {
	return "Fires rapid"
}

// gun
type Gun struct {
	power          int
	name           string
	attackStrategy iattack
}

type IGun interface {
	setName(name string)
	getName() string
	setPower(power int)
	getPower() int
	getAttack() string
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

func (g *Gun) getAttack() string {
	return g.attackStrategy.attack()
}

type ak47 struct {
	Gun
}

func newAk47(attack iattack) IGun {
	return &ak47{
		Gun{
			attackStrategy: attack,
			power:          5,
			name:           "KALASH",
		},
	}
}

type musket struct{ Gun }

func newMusket(attack iattack) IGun {
	return &musket{
		Gun{
			attackStrategy: attack,
			power:          1,
			name:           "Musket",
		},
	}
}

func getGun(name string) IGun {

	if name == "kalash" {
		return newAk47(&rapidAttack{})
	}
	if name == "musket" {
		return newMusket(&slowAttack{})
	}
	return nil
}

func main() {
	ak47 := getGun("kalash")
	printDetails(ak47)
}
func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
	fmt.Printf("Attack: %s", g.getAttack())
	fmt.Println()
	fmt.Println()
}
