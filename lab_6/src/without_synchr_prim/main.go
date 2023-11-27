package main

import (
	"fmt"
	"sync"
)

func main() {
	var pc PlayerCounter
	var wg sync.WaitGroup

	wg.Add(50)
	for i := 0; i < 50; i++ {
		go pc.Increment(&wg)
	}

	wg.Wait()
	fmt.Printf("Number of players after increment: %d\n", pc.GetCount())

	wg.Add(15)
	for i := 0; i < 15; i++ {
		go pc.Decrement(&wg)
	}

	wg.Wait()
	fmt.Printf("Number of players after decrement: %d\n", pc.GetCount())
}
