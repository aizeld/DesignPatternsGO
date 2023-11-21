package main

import "fmt"

type mac struct{}

type Computer interface{ InsertLighting() }

type Client struct{}

func (c *Client) InserIntoLighting(com Computer) {
	com.InsertLighting()
}

func (m *mac) InsertLighting() {
	fmt.Println("Lightning inserted into mac")
}

type window struct{}

func (w *window) InsertIntoUsb() {
	fmt.Println("Usb inserted into windows")
}

type WindowsAdapter struct {
	windowsMachine *window
}

func (w *WindowsAdapter) InsertLighting() {
	fmt.Println("adapting lighting into windows")
	w.windowsMachine.InsertIntoUsb()
}

func main() {
	Client := &Client{}

	// mac := &mac{}
	// Client.InserIntoLighting(mac) ----Lightning inserted into mac

	windowMachine := &window{}
	Adapter := &WindowsAdapter{windowsMachine: windowMachine}
	Client.InserIntoLighting(Adapter)

}
