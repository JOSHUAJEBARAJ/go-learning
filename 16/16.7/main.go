	package main

	import (
		"fmt"
	"sync")
	func speak(i int,g *sync.WaitGroup){
		fmt.Println(i)
		g.Done()

	}
	func main(){
		var a sync.WaitGroup
	for i:=0;i<5;i++{
		a.Add(i)
		go speak(i,&a)
	}
	a.Wait()
	fmt.Println("test")
	}