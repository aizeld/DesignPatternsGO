package main

import "fmt"

type wrestler struct {
	height int
	name   string
}

type IWrestler interface {
	setHeight(height int)
	getHeight() int
	setName(name string)
	getName() string
}

func (f *wrestler) setHeight(height int) {
	f.height = height
}

func (f *wrestler) getHeight() int {
	return f.height
}

func (f *wrestler) setName(name string) {
	f.name = name
}

func (f *wrestler) getName() string {
	return f.name
}

type khabib struct{}

type khabibWrestler struct {
	wrestler
}

type conorWrestler struct {
	wrestler
}

///striker

type striker struct {
	height int
	name   string
}

type IStriker interface {
	setHeight(height int)
	getHeight() int
	setName(name string)
	getName() string
}

func (f *striker) setHeight(height int) {
	f.height = height
}

func (f *striker) getHeight() int {
	return f.height
}

func (f *striker) setName(name string) {
	f.name = name
}

func (f *striker) getName() string {
	return f.name
}

type conor struct{}
type conorStriker struct {
	striker
}
type khabibStriker struct {
	striker
}

func (c *conor) newStriker() IStriker {
	return &conorStriker{
		striker: striker{
			name:   "conor",
			height: 175,
		},
	}
}

func (k *khabib) newStriker() IStriker {
	return &khabibStriker{
		striker: striker{
			name:   "khabib striker",
			height: 175,
		},
	}
}

func (k *conor) newWrestler() IWrestler {
	return &conorWrestler{
		wrestler: wrestler{
			name:   "conorWrestler",
			height: 178,
		},
	}
}

func (k *khabib) newWrestler() IWrestler {
	return &khabibWrestler{
		wrestler: wrestler{
			name:   "khabib",
			height: 178,
		},
	}
}

type Ifighter interface {
	newStriker() IStriker
	newWrestler() IWrestler
}

func getFighter(name string) Ifighter {
	if name == "khabib" {
		return &khabib{}
	}
	if name == "conor" {
		return &conor{}
	}
	return nil
}

///i know code looks ugly..

func main() {
	strikerConor := getFighter("conor")
	conor := strikerConor.newStriker()
	conorW := getFighter("conor")
	conorww := conorW.newWrestler()
	fmt.Println(conor.getName())
	fmt.Println(conorww.getHeight())
}
