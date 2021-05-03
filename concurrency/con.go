package main

import (
	"fmt"
	"sync"
)

func say(wg *sync.WaitGroup) {

	fmt.Println("Hello")
	wg.Done()
}
func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go say(wg)

	wg.Wait()

}
