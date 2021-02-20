package main

import(
	"fmt"
)

type distance int
type time int

func main(){
	var d distance=100
	var t time=10
	//fmt.Println(d/t) // error
	fmt.Println(time(d)/t) 

}