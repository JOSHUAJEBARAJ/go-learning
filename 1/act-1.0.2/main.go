package main

import "fmt"

func main() {

  a, b := 5, 10

  // call swap here
  swap(&a,&b)

  fmt.Println(a == 10, b == 5)

}

func swap(a *int, b *int) {

	c:=*a
	*a=*b
	*b=c
	
	//fmt.Println(*a)
// *a, *b = *b, *a ---> Simple solution
  // swap the values here

}