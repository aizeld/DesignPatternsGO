package main

import "fmt"

//Publisher
type Subject interface {
	notifyAll()
	register()
	unregister()
}

type Item struct {
	observerList []Observer
	name         string
}

func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}
func (i *Item) register(o Observer) {
	i.observerList = append(i.observerList, o)
}
func (i *Item) unregister(o Observer) {
	i.observerList = deleteFromSlice(o, i.observerList)
}

func deleteFromSlice(o Observer, observerlist []Observer) []Observer {
	length := len(observerlist)

	for i, observer := range observerlist {
		if observer.getId() == o.getId() {
			observerlist[length-1], observerlist[i] = observerlist[i], observerlist[length-1]
			return observerlist[:length-1]
		}
	}
	return observerlist
}

//Observer

type Observer interface {
	update(name string)
	getId() int
}

type Customer struct {
	id int
}

func (c *Customer) update(name string) {
	fmt.Println("customer with id:", c.id, "has received", name)
}
func (c *Customer) getId() int {
	return c.id
}

func main() {
	nike := newItem("SHIT")

	bob := &Customer{234234423}

	nike.register(bob)
	nike.notifyAll()
}
