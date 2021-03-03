package main
import (
	"fmt"
	"sync"
)
var count=0
func sum(a *sync.WaitGroup,mutex *sync.Mutex){
	mutex.Lock()
	count++
	mutex.Unlock()
	a.Done()
}
func main(){
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for i:=1;i<=500;i++{
		wg.Add(1)
		go sum(&wg,&mutex)

	}
	wg.Wait()
	fmt.Println(count)

}