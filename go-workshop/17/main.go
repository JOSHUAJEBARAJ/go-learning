package main

import "fmt"
func main(){
	m:=map[int]string{
		1:"Joshua",
		2:"Jebaraj",
	}
	if elem,ok:=m[1];ok{
		fmt.Println(elem)
	}
}