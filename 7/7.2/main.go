package main

import "fmt"

type rectangle struct{
	length int
	breadth int
}
type square struct{
	length int 
}
type area interface{
	area() int  
}

func (r rectangle)area() int{
	return r.length*r.breadth
}
func (s square)area() int{
	return s.length*s.length
}
func calculateArea(a...area){
	for _,key:=range a{
fmt.Println(key.area())
	}
}
func main(){
	r:=rectangle{10,20}
	s:=square{10}
	calculateArea(r,s)


}