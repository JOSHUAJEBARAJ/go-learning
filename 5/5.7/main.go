package main

import "fmt"

func decrement(i int) func() int{
return func() int {
	i --
	return i 
}
}

func main(){
count:=10
x:=decrement(count)
fmt.Println(x())
}