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
	v2=vector{x:1}
	v3=vector{}
	)
	fmt.Println(v1,v2)
	fmt.Println(v3.x)


}