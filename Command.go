package main

import "fmt"

type tv struct {
	isrunning bool
}

type device interface {
	TurnOff()
	TurnOn()
}

func (t *tv) TurnOff() {
	fmt.Println("Turned off bitch")
}
func (t *tv) TurnOn() {
	fmt.Println("turned on bitch")
}

//command

type button interface {
	Execute()
}

type OffCommand struct {
	device device
}

type OnCommand struct {
	device device
}

func (o *OffCommand) Execute() {
	o.device.TurnOff()
}

func (o *OnCommand) Execute() {
	o.device.TurnOn()
}

func main() {
	tv := &tv{}

	onCommand := &OffCommand{tv}
	offComand := &OnCommand{tv}

	onCommand.Execute()
	offComand.Execute()

}
