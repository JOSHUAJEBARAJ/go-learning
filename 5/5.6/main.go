package main

import "fmt"

func main(){

num:=10
x:=func (i int)int {

	return i*i
}

fmt.Printf("The sqrt of %d is %d",num,x(num))
}