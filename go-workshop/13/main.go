package main

import (
	"fmt"
)
type vector struct{
	x int
	y int
}
func main(){
	var(
	v1=vector{1,2}
	
	)
	i:=&v1
	fmt.Println(i.y)
a:=2
b:=&a
fmt.Println(b)
fmt.Println(*b)
*b=1
fmt.Println(a)

}