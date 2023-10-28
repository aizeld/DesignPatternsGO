package main

import "fmt"

// engine is receiver

type Engine struct {
	isRunning bool
}

func (engine *Engine) Start() {
	if !engine.isRunning {
		engine.isRunning = true
		fmt.Println("Started")
	} else {
		fmt.Println("Already startet")
	}
}

func (engine *Engine) Stop() {
	if engine.isRunning {
		engine.isRunning = false
		fmt.Println("Stopped")
	} else {
		fmt.Println("Already off")

	}
}

type Command interface {
	Execute()
}

type StartEngine struct {
	engine *Engine
}

func (c *StartEngine) Execute() {
	c.engine.Start()
}

type StopEngine struct {
	engine *Engine
}

func (c *StopEngine) Execute() {
	c.engine.Stop()
}

func main() {
	engine := &Engine{}
	StartCommand := &StartEngine{engine: engine}
	StopCommand := &StopEngine{engine: engine}

	StartCommand.Execute() #Started
	StopCommand.Execute() #Stopped

}
