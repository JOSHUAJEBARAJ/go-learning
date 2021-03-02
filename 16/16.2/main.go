package main

import "fmt"
import "sync"
func sum(a int ,s *sync.WaitGroup) int {
	s.Done()
	total:=0
	for i:=1 ; i<=a ;i++{
		total=i+total
	}
	return total
}
func main(){
	var s int 
	w:=&sync.WaitGroup{}
	w.Add(1)
	go func(){
		s=sum(100,w)
	}()
w.Wait()
	fmt.Println(s)

}