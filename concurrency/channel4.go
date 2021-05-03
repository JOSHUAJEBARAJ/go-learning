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

		for msg := range ch {
			fmt.Println(msg)
		}
		wg.Done()
	}(wg, ch)

	// sending

	go func(wg *sync.WaitGroup, ch chan int) {

		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}(wg, ch)

	wg.Wait()

}
