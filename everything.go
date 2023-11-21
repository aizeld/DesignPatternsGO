package main

import (
	"fmt"
	"sync"
)

type single struct{}

var lock = sync.Mutex{}

var singleInstance *single

func runSingle() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("creating instance")
			singleInstance = &single{}
		} else {
			fmt.Println("already")
		}
	} else {
		fmt.Println("already maybe")
	}
	return singleInstance
}

func main() {
	for i := 0; i < 10; i++ {
		go runSingle()
	}
	fmt.Scanln()
}
