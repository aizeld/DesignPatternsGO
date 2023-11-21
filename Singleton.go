package main

import (
	"fmt"
	"sync"
)

var lock = sync.Mutex{}

type single struct{}

var singleInstance *single

func getIntance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			fmt.Println("creating shi")
			singleInstance = &single{}
		} else {
			fmt.Println("already")
		}
	} else {
		fmt.Println("already")
	}
	return singleInstance
}

func main() {

	for i := 0; i < 10; i++ {
		go getIntance()
	}

	fmt.Scanln()
}
