package main

import (
	"fmt"
	"sync"
)

type PlayerCounter struct {
	count int
	mutex sync.Mutex
}

func (pc *PlayerCounter) Increment(wg *sync.WaitGroup) {
	pc.mutex.Lock()
	pc.count++
	fmt.Println("One more player has joined the server. Current number of players:", pc.count)
	pc.mutex.Unlock()
	wg.Done()
}

func (pc *PlayerCounter) Decrement(wg *sync.WaitGroup) {
	pc.mutex.Lock()
	pc.count--
	fmt.Println("1 player left the server. Current number of players:", pc.count)
	pc.mutex.Unlock()
	wg.Done()
}

func (pc *PlayerCounter) GetCount() int {
	pc.mutex.Lock()
	defer pc.mutex.Unlock()
	return pc.count
}
