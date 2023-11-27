package main

import (
	"fmt"
	"sync"
)

type PlayerCounter struct {
	count int
}

func (pc *PlayerCounter) Increment(wg *sync.WaitGroup) {
	pc.count++
	fmt.Println("One more player has joined the server. Current number of players:", pc.count)
	wg.Done()
}

func (pc *PlayerCounter) Decrement(wg *sync.WaitGroup) {
	pc.count--
	fmt.Println("1 player left the server. Current number of players", pc.count)
	wg.Done()
}

func (pc *PlayerCounter) GetCount() int {
	return pc.count
}
