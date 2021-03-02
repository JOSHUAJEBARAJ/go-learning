package main

import "fmt"
func sum(a int ) int {
	total:=0
	for i:=1 ; i<=a ;i++{
		total=i+total
	}
	return total
}
func main(){
	var s int 
	go func(){
		s=sum(100)
	}()
	b:=sum(10)
	fmt.Println(s,b)

}