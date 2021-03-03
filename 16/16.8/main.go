package main
import (
	"fmt"
	"sync"
)
var count=0
func sum(a *sync.WaitGroup){
	count++
	a.Done()
}
func main(){
	var wg sync.WaitGroup
	for i:=0;i<=500;i++{
		wg.Add(1)
		go sum(&wg)

	}
	wg.Wait()
	fmt.Println(count)

}