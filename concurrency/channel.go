package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	// recieveing

	wg.Add(2)
	go func(wg *sync.WaitGroup, ch chan int) {

		fmt.Println(<-ch)
		wg.Done()
	}(wg, ch)

	// sending

	go func(wg *sync.WaitGroup, ch chan int) {

		ch <- 42
		ch <- 2
		wg.Done()
	}(wg, ch)

	wg.Wait()

}
