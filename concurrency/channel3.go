package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	// recieveing

	wg.Add(2)
	go func(wg *sync.WaitGroup, ch chan int) {

		msg, ok := <-ch
		if ok {
			fmt.Println(msg, ok)
		}
		wg.Done()
	}(wg, ch)

	// sending

	go func(wg *sync.WaitGroup, ch chan int) {

		close(ch)

		wg.Done()
	}(wg, ch)

	wg.Wait()

}
