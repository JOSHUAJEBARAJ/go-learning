package main

import "fmt"

func sum(i ...int)int {

	total:=0
	for _,value := range i{
total+=value
	}

	return total 
}
func main(){
	a:=sum(10,20)
fmt.Println(a)
}